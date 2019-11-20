package message


import (
	"errors"
	"io"
)

type Buffer struct {
	reader io.Reader
	buf    []byte
	start  int
	end    int
	len    int
}

func NewBuffer(reader io.Reader, len int) Buffer {
	buf := make([]byte, len)
	return Buffer{reader, buf, 0, 0,len}
}

func (b *Buffer) Len() int {
	return b.end - b.start
}

//将有用的字节前移
func (b *Buffer) grow() {
	if b.start == 0 {
		if b.end-b.start == len(b.buf){
			buf := make([]byte, len(b.buf)+b.len)
			copy(buf,b.buf)
			b.buf = buf
		}
		return
	}
	copy(b.buf, b.buf[b.start:b.end])
	b.end -= b.start
	b.start = 0
	if b.end-b.start == len(b.buf){
		 buf := make([]byte, len(b.buf)+b.len)
		 copy(buf,b.buf)
		 b.buf = buf
	}
}

//从reader里面读取数据，如果reader阻塞，会发生阻塞
func (b *Buffer) ReadFromReader() (int, error) {
	b.grow()
	n, err := b.reader.Read(b.buf[b.end:])
	if err != nil {
		return n, err
	}
	b.end += n
	return n, nil
}

//返回n个字节，而不产生移位
func (b *Buffer) Seek(n int) ([]byte, error) {
	if b.end-b.start >= n {
		buf := b.buf[b.start:b.start+n]
		return buf, nil
	}
	return nil, errors.New("not enough")
}

func (b *Buffer)  SeekOffset(offset int,n int) ([]byte, error) {
	if b.end-b.start - offset>= n {
		buf := b.buf[b.start+offset:b.start+offset+n]
		return buf, nil
	}
	return nil, errors.New("not enough")
}

//舍弃offset个字段，读取n个字段
func (b *Buffer) Read(offset, n int) ([]byte) {
	b.start += offset
	buf := b.buf[b.start:b.start+n]
	b.start += n
	return buf
}