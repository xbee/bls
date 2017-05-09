package bls

/*
#cgo CFLAGS:-I../../include -DBLS_MAX_OP_UNIT_SIZE=6
#cgo bn256 CFLAGS:-UBLS_MAX_OP_UNIT_SIZE -DBLS_MAX_OP_UNIT_SIZE=4
#cgo bn384 CFLAGS:-UBLS_MAX_OP_UNIT_SIZE -DBLS_MAX_OP_UNIT_SIZE=6
#cgo LDFLAGS:-lbls -lbls_if -lmcl -lgmp -lgmpxx -L../lib -L../../lib -L../../../mcl/lib -L../../mcl/lib  -lstdc++ -lcrypto
#include "bls_if.h"
*/
import "C"
import "fmt"
import "unsafe"
import "encoding/hex"

// CurveFp254BNb -- 254 bit curve
const CurveFp254BNb = 0

// CurveFp382_1 -- 382 bit curve 1
const CurveFp382_1 = 1

// CurveFp382_2 -- 382 bit curve 2
const CurveFp382_2 = 2

// Init --
// call this function before calling all the other operations
// this function is not thread safe
func Init(curve int) {
	C.blsInit(C.int(curve), C.BLS_MAX_OP_UNIT_SIZE)
}

// GetMaxOpUnitSize --
func GetMaxOpUnitSize() int {
	return int(C.BLS_MAX_OP_UNIT_SIZE)
}

// GetOpUnitSize --
func GetOpUnitSize() int {
	return int(C.blsGetOpUnitSize())
}

// GetCurveOrder --
func GetCurveOrder() string {
	buf := make([]byte, 1024)
	// #nosec
	n := C.blsGetCurveOrder((*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)))
	if n == 0 {
		panic("implementation err. size of buf is small")
	}
	return string(buf[:n])
}

// GetFieldOrder --
func GetFieldOrder() string {
	buf := make([]byte, 1024)
	// #nosec
	n := C.blsGetFieldOrder((*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)))
	if n == 0 {
		panic("implementation err. size of buf is small")
	}
	return string(buf[:n])
}

// ID --
type ID struct {
	v [C.BLS_MAX_OP_UNIT_SIZE]C.uint64_t
}

// getPointer --
func (id *ID) getPointer() (p *C.blsId) {
	// #nosec
	return (*C.blsId)(unsafe.Pointer(&id.v[0]))
}

// GetByte
func (id *ID) GetByte(ioMode int) []byte {
	buf := make([]byte, 1024)
	// #nosec
	n := C.blsIdGetStr(id.getPointer(), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), C.int(ioMode))
	if n == 0 {
		panic("implementation err. size of buf is small")
	}
	return buf[:n]
}

// SetByte --
func (id *ID) SetByte(buf []byte, ioMode int) error {
	// #nosec
	err := C.blsIdSetStr(id.getPointer(), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), C.int(ioMode))
	if err > 0 {
		return fmt.Errorf("bad byte:%x", buf)
	}
	return nil
}

// Serialize
func (id *ID) Serialize() []byte {
	return id.GetByte(C.blsIoEcComp)
}

// Deserialize
func (id *ID) Deserialize(b []byte) error {
	return id.SetByte(b, C.blsIoEcComp)
}

// GetHexString
func (id *ID) GetHexString() string {
	return string(id.GetByte(16))
}

// GetDecString
func (id *ID) GetDecString() string {
	return string(id.GetByte(10))
}

// SetHexString
func (id *ID) SetHexString(s string) error {
	return id.SetByte([]byte(s), 16)
}

// SetDecString
func (id *ID) SetDecString(s string) error {
	return id.SetByte([]byte(s), 10)
}

// IsSame --
func (id *ID) IsSame(rhs *ID) bool {
	return C.blsIdIsSame(id.getPointer(), rhs.getPointer()) == 1
}

// Set --
func (id *ID) Set(v []uint64) {
	expect := GetOpUnitSize()
	if len(v) != expect {
		panic(fmt.Errorf("bad size (%d), expected size %d", len(v), expect))
	}
	// #nosec
	C.blsIdSet(id.getPointer(), (*C.uint64_t)(unsafe.Pointer(&v[0])))
}

// SecretKey --
type SecretKey struct {
	v [C.BLS_MAX_OP_UNIT_SIZE]C.uint64_t
}

// getPointer --
func (sec *SecretKey) getPointer() (p *C.blsSecretKey) {
	// #nosec
	return (*C.blsSecretKey)(unsafe.Pointer(&sec.v[0]))
}

// GetByte
func (sec *SecretKey) GetByte(ioMode int) []byte {
	buf := make([]byte, 1024)
	// #nosec
	n := C.blsSecretKeyGetStr(sec.getPointer(), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), C.int(ioMode))
	if n == 0 {
		panic("implementation err. size of buf is small")
	}
	return buf[:n]
}

// SetByte --
func (sec *SecretKey) SetByte(buf []byte, ioMode int) error {
	// #nosec
	err := C.blsSecretKeySetStr(sec.getPointer(), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), C.int(ioMode))
	if err > 0 {
		return fmt.Errorf("bad byte:%x", buf)
	}
	return nil
}

// Serialize
func (sec *SecretKey) Serialize() []byte {
	return sec.GetByte(C.blsIoEcComp)
}

// Deserialize
func (sec *SecretKey) Deserialize(b []byte) error {
	return sec.SetByte(b, C.blsIoEcComp)
}

// GetHexString
func (sec *SecretKey) GetHexString() string {
	return string(sec.GetByte(16))
}

// GetDecString
func (sec *SecretKey) GetDecString() string {
	return string(sec.GetByte(10))
}

// SetHexString
func (sec *SecretKey) SetHexString(s string) error {
	return sec.SetByte([]byte(s), 16)
}

// SetDecString
func (sec *SecretKey) SetDecString(s string) error {
	return sec.SetByte([]byte(s), 10)
}

// IsSame --
func (sec *SecretKey) IsSame(rhs *SecretKey) bool {
	return C.blsSecretKeyIsSame(sec.getPointer(), rhs.getPointer()) == 1
}

// SetArray --
func (sec *SecretKey) SetArray(v []uint64) {
	expect := GetOpUnitSize()
	if len(v) != expect {
		panic(fmt.Errorf("bad size (%d), expected size %d", len(v), expect))
	}
	// #nosec
	C.blsSecretKeySetArray(sec.getPointer(), (*C.uint64_t)(unsafe.Pointer(&v[0])))
}

// Init --
func (sec *SecretKey) Init() {
	C.blsSecretKeyInit(sec.getPointer())
}

// Add --
func (sec *SecretKey) Add(rhs *SecretKey) {
	C.blsSecretKeyAdd(sec.getPointer(), rhs.getPointer())
}

// GetMasterSecretKey --
func (sec *SecretKey) GetMasterSecretKey(k int) (msk []SecretKey) {
	msk = make([]SecretKey, k)
	msk[0] = *sec
	for i := 1; i < k; i++ {
		msk[i].Init()
	}
	return msk
}

// GetMasterPublicKey --
func GetMasterPublicKey(msk []SecretKey) (mpk []PublicKey) {
	n := len(msk)
	mpk = make([]PublicKey, n)
	for i := 0; i < n; i++ {
		mpk[i] = *msk[i].GetPublicKey()
	}
	return mpk
}

// Set --
func (sec *SecretKey) Set(msk []SecretKey, id *ID) {
	C.blsSecretKeySet(sec.getPointer(), msk[0].getPointer(), C.size_t(len(msk)), id.getPointer())
}

// Recover --
func (sec *SecretKey) Recover(secVec []SecretKey, idVec []ID) {
	C.blsSecretKeyRecover(sec.getPointer(), secVec[0].getPointer(), idVec[0].getPointer(), C.size_t(len(secVec)))
}

// GetPop --
func (sec *SecretKey) GetPop() (sign *Sign) {
	sign = new(Sign)
	C.blsSecretKeyGetPop(sec.getPointer(), sign.getPointer())
	return sign
}

// PublicKey --
type PublicKey struct {
	v [C.BLS_MAX_OP_UNIT_SIZE * 2 * 3]C.uint64_t
}

// getPointer --
func (pub *PublicKey) getPointer() (p *C.blsPublicKey) {
	// #nosec
	return (*C.blsPublicKey)(unsafe.Pointer(&pub.v[0]))
}

// GetByte
func (pub *PublicKey) GetByte(ioMode int) []byte {
	buf := make([]byte, 1024)
	// #nopub
	n := C.blsPublicKeyGetStr(pub.getPointer(), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), C.int(ioMode))
	if n == 0 {
		panic("implementation err. size of buf is small")
	}
	return buf[:n]
}

// SetByte --
func (pub *PublicKey) SetByte(buf []byte, ioMode int) error {
	// #nopub
	err := C.blsPublicKeySetStr(pub.getPointer(), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), C.int(ioMode))
	if err > 0 {
		return fmt.Errorf("bad byte:%x", buf)
	}
	return nil
}

// Serialize
func (pub *PublicKey) Serialize() []byte {
	return pub.GetByte(C.blsIoEcComp)
}

// Deserialize
func (pub *PublicKey) Deserialize(b []byte) error {
	return pub.SetByte(b, C.blsIoEcComp)
}

// GetHexString
func (pub *PublicKey) GetHexString() string {
	return fmt.Sprintf("%x", pub.Serialize())
}

// SetStr
func (pub *PublicKey) SetStr(s string) error {
	b, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	return pub.Deserialize(b)
}

// IsSame --
func (pub *PublicKey) IsSame(rhs *PublicKey) bool {
	return C.blsPublicKeyIsSame(pub.getPointer(), rhs.getPointer()) == 1
}

// Add --
func (pub *PublicKey) Add(rhs *PublicKey) {
	C.blsPublicKeyAdd(pub.getPointer(), rhs.getPointer())
}

// Set --
func (pub *PublicKey) Set(mpk []PublicKey, id *ID) {
	C.blsPublicKeySet(pub.getPointer(), mpk[0].getPointer(), C.size_t(len(mpk)), id.getPointer())
}

// Recover --
func (pub *PublicKey) Recover(pubVec []PublicKey, idVec []ID) {
	C.blsPublicKeyRecover(pub.getPointer(), pubVec[0].getPointer(), idVec[0].getPointer(), C.size_t(len(pubVec)))
}

// Sign  --
type Sign struct {
	v [C.BLS_MAX_OP_UNIT_SIZE * 3]C.uint64_t
}

// getPointer --
func (sign *Sign) getPointer() (p *C.blsSign) {
	// #nosec
	return (*C.blsSign)(unsafe.Pointer(&sign.v[0]))
}

// GetByte
func (sign *Sign) GetByte(ioMode int) []byte {
	buf := make([]byte, 1024)
	// #nosign
	n := C.blsSignGetStr(sign.getPointer(), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), C.int(ioMode))
	if n == 0 {
		panic("implementation err. size of buf is small")
	}
	return buf[:n]
}

// SetByte --
func (sign *Sign) SetByte(buf []byte, ioMode int) error {
	// #nosign
	err := C.blsSignSetStr(sign.getPointer(), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), C.int(ioMode))
	if err > 0 {
		return fmt.Errorf("bad byte:%x", buf)
	}
	return nil
}

// Serialize
func (sign *Sign) Serialize() []byte {
	return sign.GetByte(C.blsIoEcComp)
}

// Deserialize
func (sign *Sign) Deserialize(b []byte) error {
	return sign.SetByte(b, C.blsIoEcComp)
}

// GetHexString
func (sign *Sign) GetHexString() string {
	return fmt.Sprintf("%x", sign.Serialize())
}

// SetStr
func (sign *Sign) SetStr(s string) error {
	b, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	return sign.Deserialize(b)
}

// IsSame --
func (sign *Sign) IsSame(rhs *Sign) bool {
	return C.blsSignIsSame(sign.getPointer(), rhs.getPointer()) == 1
}

// GetPublicKey --
func (sec *SecretKey) GetPublicKey() (pub *PublicKey) {
	pub = new(PublicKey)
	C.blsSecretKeyGetPublicKey(sec.getPointer(), pub.getPointer())
	return pub
}

// Sign -- Constant Time version
func (sec *SecretKey) Sign(m string) (sign *Sign) {
	sign = new(Sign)
	buf := []byte(m)
	// #nosec
	C.blsSecretKeySign(sec.getPointer(), sign.getPointer(), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)))
	return sign
}

// Add --
func (sign *Sign) Add(rhs *Sign) {
	C.blsSignAdd(sign.getPointer(), rhs.getPointer())
}

// Recover --
func (sign *Sign) Recover(signVec []Sign, idVec []ID) {
	C.blsSignRecover(sign.getPointer(), signVec[0].getPointer(), idVec[0].getPointer(), C.size_t(len(signVec)))
}

// Verify --
func (sign *Sign) Verify(pub *PublicKey, m string) bool {
	buf := []byte(m)
	// #nosec
	return C.blsSignVerify(sign.getPointer(), pub.getPointer(), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf))) == 1
}

// VerifyPop --
func (sign *Sign) VerifyPop(pub *PublicKey) bool {
	return C.blsSignVerifyPop(sign.getPointer(), pub.getPointer()) == 1
}