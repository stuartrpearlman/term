// +build !windows

package termios

import (
	"syscall"
	"unsafe"
)

// Tiocmget returns the state of the MODEM bits.
func Tiocmget(fd uintptr, status *int) error {
	return ioctl(fd, syscall.TIOCMGET, uintptr(unsafe.Pointer(status)))
}

// Tiocmset sets the state of the MODEM bits.
func Tiocmset(fd uintptr, status *int) error {
	return ioctl(fd, syscall.TIOCMSET, uintptr(unsafe.Pointer(status)))
}

// Tiocmbis sets the indicated modem bits.
func Tiocmbis(fd uintptr, status *int) error {
	return ioctl(fd, syscall.TIOCMBIS, uintptr(unsafe.Pointer(status)))
}

// Tiocmbic clears the indicated modem bits.
func Tiocmbic(fd uintptr, status *int) error {
	return ioctl(fd, syscall.TIOCMBIC, uintptr(unsafe.Pointer(status)))
}

// Cfmakecbreak modifies attr for cbreak mode.
func Cfmakecbreak(attr *syscall.Termios) {
	attr.Lflag &^= syscall.ECHO | syscall.ICANON
	attr.Cc[syscall.VMIN] = 1
	attr.Cc[syscall.VTIME] = 0
}

// Cfmakeraw modifies attr for raw mode.
func Cfmakeraw(attr *syscall.Termios) {
	attr.Iflag &^= syscall.BRKINT | syscall.ICRNL | syscall.INPCK | syscall.ISTRIP | syscall.IXON
	attr.Oflag &^= syscall.OPOST
	attr.Cflag &^= syscall.CSIZE | syscall.PARENB
	attr.Cflag |= syscall.CS8
	attr.Lflag &^= syscall.ECHO | syscall.ICANON | syscall.IEXTEN | syscall.ISIG
	attr.Cc[syscall.VMIN] = 1
	attr.Cc[syscall.VTIME] = 0
}

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

// Tiocgwinsz returns the window size.
func Tiocgwinsz(fd uintptr) (*Winsize, error) {
	var ws Winsize
	err := ioctl(fd, syscall.TIOCGWINSZ, uintptr(unsafe.Pointer(&ws)))
	return &ws, err
}

// Tiocswinsz sets the window size.
func Tiocswinsz(fd uintptr, value *Winsize) error {
	return ioctl(fd, syscall.TIOCSWINSZ, uintptr(unsafe.Pointer(value)))
}
