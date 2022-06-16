package etxhash

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const NODE_KEY = "8W7QlFdpnjaG3i72"
const USE_NODE_KEY = true

func AESBase64Encrypt(origin_data string, key string) (base64_result string, err error) {
	iv := []byte(key) //"1234567890123456"
	var block cipher.Block
	if block, err = aes.NewCipher([]byte(key)); err != nil {
		log.Println(err)
		return
	}

	encrypt := cipher.NewCBCEncrypter(block, iv)
	var source []byte = PKCS5Padding([]byte(origin_data), 16)
	var dst []byte = make([]byte, len(source))
	encrypt.CryptBlocks(dst, source)
	base64_result = base64.StdEncoding.EncodeToString(dst)
	return
}

func AESBase64Decrypt(encrypt_data string, key string) (origin_data string, err error) {
	iv := []byte(key)
	var block cipher.Block
	if block, err = aes.NewCipher([]byte(key)); err != nil {
		log.Println(err)
		return
	}
	encrypt := cipher.NewCBCDecrypter(block, iv)

	var source []byte
	if source, err = base64.StdEncoding.DecodeString(encrypt_data); err != nil {
		log.Println(err)
		return
	}
	var dst []byte = make([]byte, len(source))
	encrypt.CryptBlocks(dst, source)
	origin_data = string(PKCS5Unpadding(dst))
	return
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Unpadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func ActualHash(s string) int32 {
	var h int32 = 0
	for i := 0; i < len(s); i++ {
		c1 := s[i]
		v1 := rune(c1)
		h = 31*h + v1
	}
	return h
}

func GetSignature(requestTimeMills string, accountName string, ownerKey string, activeKey string) string {
	var linkStrings = []string{requestTimeMills, accountName, ownerKey, "mtc", activeKey}
	linkStr := strings.Join(linkStrings, "&")
	fmt.Println(linkStr)
	hash_value := ActualHash(linkStr)
	fmt.Println("Result: ", hash_value)
	hash_str := strconv.Itoa(int(hash_value))

	//	crypto.MD5.New()
	//return CodecUtil.getMD5Code(String.valueOf(linkStr.hashCode()));
	md5Hasher := md5.New()
	md5Hasher.Write([]byte(hash_str))
	md5Str := hex.EncodeToString(md5Hasher.Sum(nil))
	fmt.Println(md5Str)
	return md5Str
}

func DesBase64Encrypt(origData, key []byte) (result string, err error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	base64_result := base64.StdEncoding.EncodeToString(crypted)
	return base64_result, nil
}

func DesBase64Decrypt(crypted, key []byte) (result string, err error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5Unpadding(origData)
	// origData = ZeroUnPadding(origData)
	base64_result := base64.StdEncoding.EncodeToString(origData)
	return base64_result, nil
}

func MakeAESKey(requestTimeMills int64, singnature string) string {
	temp := strconv.FormatInt(requestTimeMills-2, 10) + singnature
	hash_value := ActualHash(temp)
	hash_string := strconv.FormatInt(int64(hash_value), 10)
	md5Hasher := md5.New()
	md5Hasher.Write([]byte(hash_string))
	md5Str := hex.EncodeToString(md5Hasher.Sum(nil))
	return md5Str[0:16]
}

//public static String makeKey(Long requestTimeMills, String singnature) {
//String temp = requestTimeMills - 2 + singnature;
//String md5Code = CodecUtil.getMD5Code(String.valueOf(temp.hashCode()));
//return md5Code.substring(0, 16);
//}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateSecertKey(user, password, email, phone string) (string, error) {
	//token := jwt.New(jwt.SigningMethodHS256)
	//
	//// Set claims
	//claims := token.Claims.(jwt.MapClaims)
	//claims["user"] = user
	//claims["password"] = password
	//claims["email"] = email
	//claims["phone"] = phone
	//
	//// Generate encoded token and send it as response.
	//t, err := token.SignedString([]byte("H6qerrGqAcNXafHZ"))
	//fmt.Println(t, err)
	var linkStrings = []string{user, password, email, "mtc", phone}
	linkStr := strings.Join(linkStrings, "&")
	fmt.Println(linkStr)
	hash_value := ActualHash(linkStr)
	fmt.Println("Result: ", hash_value)
	hash_str := strconv.Itoa(int(hash_value))

	//	crypto.MD5.New()
	//return CodecUtil.getMD5Code(String.valueOf(linkStr.hashCode()));
	md5Hasher := md5.New()
	md5Hasher.Write([]byte(hash_str))
	md5Str := hex.EncodeToString(md5Hasher.Sum(nil))
	fmt.Println(md5Str)
	return md5Str, nil
}

func VerifyAIPFundPassword(psw1, psw2 string) bool {
	var h hash.Hash = sha1.New()
	h.Write([]byte(psw1))
	bs := h.Sum(nil)
	//fmt.Println(psw1, psw2, bs)
	psw1Enc := base64.StdEncoding.EncodeToString([]byte(bs))
	if strings.Compare(psw1Enc, psw2) == 0 {
		return true
	}

	return false
}
