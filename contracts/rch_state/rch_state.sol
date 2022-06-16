pragma solidity >=0.4.22 <0.7.0;

contract token { function transfer(address receiver, uint amount) public pure { receiver; amount; } }

library SafeMath {

    function mul(uint256 a, uint256 b)
    internal
    pure
    returns (uint256 c)
    {
        if (a == 0) {
            return 0;
        }
        c = a * b;
        require(c / a == b, "SafeMath mul failed");
        return c;
    }

    function div(uint256 a, uint256 b)
    internal
    pure
    returns (uint256)
    {
        // assert(b > 0); // Solidity automatically throws when dividing by 0
        uint256 c = a / b;
        // assert(a == b * c + a % b); // There is no case in which this doesn't hold
        return c;
    }

    /**
    * @dev Subtracts two numbers, throws on overflow (i.e. if subtrahend is greater than minuend).
    */
    function sub(uint256 a, uint256 b)
    internal
    pure
    returns (uint256)
    {
        require(b <= a, "SafeMath sub failed");
        return a - b;
    }

    /**
    * @dev Adds two numbers, throws on overflow.
    */
    function add(uint256 a, uint256 b)
    internal
    pure
    returns (uint256 c)
    {
        c = a + b;
        require(c >= a, "SafeMath add failed");
        return c;
    }

    /**
     * @dev gives square root of given x.
     */
    function sqrt(uint256 x)
    internal
    pure
    returns (uint256 y)
    {
        uint256 z = ((add(x,1)) / 2);
        y = x;
        while (z < y)
        {
            y = z;
            z = ((add((x / z),z)) / 2);
        }
    }
    //10000×(1-5%)^180≈0.98
    /**
     * @dev gives square. multiplies x by x
     */
    function sq(uint256 x)
    internal
    pure
    returns (uint256)
    {
        return (mul(x,x));
    }

    /**
     * @dev x to the power of y
     */
    function pwr(uint256 x, uint256 y)
    internal
    pure
    returns (uint256)
    {
        if (x==0)
            return (0);
        else if (y==0)
            return (1);
        else
        {
            uint256 z = x;
            for (uint256 i=1; i < y; i++)
                z = mul(z,x);
            return (z);
        }
    }
}


contract RchPoolMgr {

    address owner;
    string public version;

    struct PoolValue {
        address userAddr ;
        uint256 userBalane;
        uint256 updateTime ;
        bool    unstakeing;  // false  staking ; true unstaking ;
        uint256 unstakeGate; //
    }

    struct BindingInfo {
        address userAddr ;
        string  devCode ; // device code ;
        bool    status;  //
    }

    mapping (address => PoolValue) public poolStakeList;
    //mapping (address => BindingInfo) public bindingList;

    mapping(address=>bool) public adminList ;

    address[]  public poolAddrList ;
    uint256    public poolListCount ;
    // event EventEtxMutilTransfer(address payable[] addrList, uint256[] amountList);
    event Transfer(address indexed _from, address indexed _to, uint indexed value);

    uint256 constant Limit10000 = (10000*10**uint256(18));
    uint256 constant unstake_time = 300 ; // 3600*24*30
    uint256 constant ETHV_REDUCE_BOUNDING = 1110858 ;  // 半年减半机制

    constructor() public {
        owner = msg.sender;
        version = "RchPoolMgr version 1.5";
    }


    function() external payable {

        revert();
    }


    function ChangeOwner(address owner2) public {
        require(msg.sender == owner, "only owner can do it !" );
        owner = owner2;
    }

    function SetAdmin(address addr, bool flag) public {
        require(msg.sender == owner , "only owner can do it !");
        adminList[addr] = flag ;
    }


    function IsAdmin(address addr ) public view returns (bool ret ) {

        return adminList[addr];
    }



    function GetStakeGateValue() public view returns (uint256 ret)  {
        //ETHV_REDUCE_BOUNDING = 1110858  // 半年减半机制
        uint256 redu = SafeMath.div(block.number, ETHV_REDUCE_BOUNDING);
        ret = SafeMath.mul(Limit10000, SafeMath.pwr(2, redu));
        return ret;
    }


    function GetStakeGateValueTest(uint256 blocknum) public pure returns (uint256 ret)
    {
        //ETHV_REDUCE_BOUNDING = 1110858  // 半年减半机制
        uint256 redu = SafeMath.div(blocknum, ETHV_REDUCE_BOUNDING);
        ret = SafeMath.mul(Limit10000, SafeMath.pwr(2, redu));
        return ret;
    }



    function RchStakePoolValue()external payable  {

        require(msg.value >= GetStakeGateValue() , "need 1000 rch for create token.");
        require(poolStakeList[msg.sender].unstakeing == false , "need 1000 rch for create token.");
        poolStakeList[msg.sender].userAddr = msg.sender ;
        poolStakeList[msg.sender].userBalane += msg.value ;
        poolStakeList[msg.sender].updateTime = now;
        //poolStakeList[msg.sender].unstakeGate = poolStakeList[msg.sender].userBalane/100;
        bool flag1 = false ;

        for(uint i= 0; i< poolAddrList.length ; i++ ){
            if(poolAddrList[i] == msg.sender){
                flag1 = true ;
                break;
            }
        }

        if(flag1 == false ){
            poolAddrList.push(msg.sender);
            poolListCount = poolListCount +1 ;
        }

    }


    function RchUnStakePoolValue()external payable  {

        if( poolStakeList[msg.sender].userBalane > 0 ){
            if( poolStakeList[msg.sender].unstakeing == false ){
                //poolStakeList[msg.sender].unstakeGate = poolStakeList[msg.sender].userBalane/100;
                //poolStakeList[msg.sender].unstakeing = true ;
            }

            require( (now - poolStakeList[msg.sender].updateTime ) > unstake_time , "It will take 30 days to unstake.");
            //require( poolStakeList[msg.sender].unstakeGate> 0 , "need unstake gate big than 0 .");
            msg.sender.transfer( poolStakeList[msg.sender].userBalane );
            emit Transfer( address(this) , msg.sender, poolStakeList[msg.sender].userBalane);

            poolStakeList[msg.sender].userBalane  = 0;
            poolStakeList[msg.sender].updateTime = now;


        }
    }


    function kill() public payable{
        if (msg.sender == owner){
            selfdestruct(address(uint160(0x7A620f1Ad467487Dd6b6353224F3fF7f5771A248)));
        }
    }


    // function EtxGetPoolValue(address userAddr)public pure returns (PoolValue memvory ret)  {
    //      PoolValue memory pv = poolStakeList[userAddr];
    //      return pv ;
    // }

}


contract RchDeviceMgr {

    address owner;
    string public version;


    mapping(bytes32=>bool) public devcodeList ;
    bytes32[]  public devcodeArr ;
    uint256    public devcodeListCount ;

    mapping(address=>bool) public adminList ;


    constructor() public {
        owner = msg.sender;
        version = "RchDeviceMgr version 1.4";
    }


    function() external payable {

        revert();
    }


    function ChangeOwner(address owner2) public {
        require(msg.sender == owner, "only owner can do it !" );
        owner = owner2;
    }


    function SetAdmin(address addr, bool flag) public {
        require(msg.sender == owner , "only owner can do it !");
        adminList[addr] = flag ;
    }


    function IsAdmin(address addr ) public view returns (bool ret ) {

        return adminList[addr];
    }


    function ImportDeviceCode(bytes32 devcode , bool status ) public {

        require(adminList[msg.sender], "only admin do it !");

        devcodeList[devcode] = status ;
        devcodeListCount = devcodeListCount +1 ;
        devcodeArr.push(devcode );

    }

    function MultiImportDeviceCode(bytes32[] memory devArr , bool[] memory statusArr ) public {

        require(adminList[msg.sender], "only admin do it !");

        for(uint i= 0; i< devArr.length; i++ ){
            bytes32 devcode = devArr[i];
            devcodeList[devcode] = statusArr[i] ;
            devcodeArr.push(devArr[i] );
            devcodeListCount = devcodeListCount +1 ;
        }

    }



    function GetDeviceCode(bytes32 devcode ) public view returns (bool success)  {

        return devcodeList[devcode];

    }


    function kill() public payable{
        if (msg.sender == owner){
            selfdestruct(address(uint160(0x7A620f1Ad467487Dd6b6353224F3fF7f5771A248)));
        }
    }


}