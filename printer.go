package zpl_printer

import (
	"errors"
	"fmt"
	"unsafe"
)

type Printer struct {
	id         uintptr
	portOpened bool
	sdk        *Sdk
}

// 定义一个枚举类型
const (
	ErrSdk = iota
	ErrPrt
	ErrPort
)

func NewPrinter(sdk *Sdk, model, portSetting string) (p *Printer, err error) {
	if sdk == nil || sdk.Proc == nil {
		err = fmt.Errorf("printer sdk invalid")
		return
	}
	p = &Printer{sdk: sdk}
	p.id, _, _ = sdk.Proc.InitPrinter.Call(String2UintPTR(model))
	if p.id == 0 {
		err = fmt.Errorf("init printer failed")
		return
	}
	err = p.OpenPort(portSetting)
	if err != nil {
		_ = p.Release()
		return nil, err
	}
	return p, nil
}

func (rec *Printer) HasError(level int) (err error) {
	if rec.sdk == nil || rec.sdk.Proc == nil {
		err = fmt.Errorf("printer sdk invalid")
		return
	}
	if level > 0 {
		if rec.id == 0 {
			err = fmt.Errorf("printer invalid")
			return
		}
	}
	if level > 1 {
		if !rec.portOpened {
			err = fmt.Errorf("printer communication port closed")
			return
		}
	}
	return
}

func (rec *Printer) Release() (err error) {
	if rec.id != 0 {
		if err = rec.HasError(ErrSdk); err != nil {
			return
		}
		ret, _, _ := rec.sdk.Proc.ReleasePrinter.Call(rec.id)
		if ret != 0 {
			err = fmt.Errorf("release printer failed,err code:%v", int32(ret))
			return
		}
		rec.id = 0
		rec.portOpened = false
		rec.sdk = nil
	}
	return
}

func (rec *Printer) OpenPort(portSetting string) (err error) {
	if err = rec.HasError(ErrPrt); err != nil {
		return
	}
	if rec.portOpened {
		err = errors.New("printer communication port already open")
		return
	}
	ret, _, _ := rec.sdk.Proc.OpenPort.Call(rec.id, String2UintPTR(portSetting))
	if ret != 0 {
		err = fmt.Errorf("open port failed,err code:%v", int32(ret))
		return
	}
	rec.portOpened = true
	return
}

func (rec *Printer) ClosePort() (err error) {
	if err = rec.HasError(ErrPort); err != nil {
		return
	}
	ret, _, _ := rec.sdk.Proc.ClosePort.Call(rec.id)
	if ret != 0 {
		err = fmt.Errorf("close port failed,err code:%v", int32(ret))
		return
	}
	rec.portOpened = false
	return
}

func (rec *Printer) StartFormat() (err error) {
	if err = rec.HasError(ErrPort); err != nil {
		return
	}
	ret, _, _ := rec.sdk.Proc.StartFormat.Call(rec.id)
	if ret != 0 {
		err = fmt.Errorf("start format failed,err code:%v", int32(ret))
		return
	}
	return
}

func (rec *Printer) EndFormat() (err error) {
	if err = rec.HasError(ErrPort); err != nil {
		return
	}
	ret, _, _ := rec.sdk.Proc.EndFormat.Call(rec.id)
	if ret != 0 {
		err = fmt.Errorf("end format failed,err code:%v", int32(ret))
		return
	}
	return
}

func (rec *Printer) JoinText(xPos, yPos, fontNum, orientation, fontWidth, fontHeight int, text string) (err error) {
	if err = rec.HasError(ErrPort); err != nil {
		return
	}
	ret, _, _ := rec.sdk.Proc.Text.Call(rec.id,
		uintptr(xPos),
		uintptr(yPos),
		uintptr(fontNum),
		uintptr(orientation),
		uintptr(fontWidth),
		uintptr(fontHeight),
		String2UintPTR(text),
	)
	if ret != 0 {
		err = fmt.Errorf("join text failed,err code:%v", int32(ret))
		return
	}
	return
}

func (rec *Printer) JoinTextBlock(xPos, yPos, fontNum, orientation, fontWidth, fontHeight, textBlockWidth, textBlockHeight int, text string) (err error) {
	if err = rec.HasError(ErrPort); err != nil {
		return
	}
	ret, _, _ := rec.sdk.Proc.TextBlock.Call(rec.id,
		uintptr(xPos),
		uintptr(yPos),
		uintptr(fontNum),
		uintptr(orientation),
		uintptr(fontWidth),
		uintptr(fontHeight),
		uintptr(textBlockWidth),
		uintptr(textBlockHeight),
		String2UintPTR(text),
	)
	if ret != 0 {
		err = fmt.Errorf("join text block failed,err code:%v", int32(ret))
		return
	}
	return
}

func (rec *Printer) JoinBarCode128(xPos, yPos, orientation, moduleWidth, codeHeight int, line, lineAboveCode, checkDigit, mode, text string) (err error) {
	if err = rec.HasError(ErrPort); err != nil {
		return
	}
	ret, _, _ := rec.sdk.Proc.BarCode128.Call(rec.id,
		uintptr(xPos),
		uintptr(yPos),
		uintptr(orientation),
		uintptr(moduleWidth),
		uintptr(codeHeight),
		String2UintPTR(line),
		String2UintPTR(lineAboveCode),
		String2UintPTR(checkDigit),
		String2UintPTR(mode),
		String2UintPTR(text),
	)
	if ret != 0 {
		err = fmt.Errorf("join barcode128 failed,err code:%v", int32(ret))
		return
	}
	return
}

func (rec *Printer) JoinQRCode(xPos, yPos, orientation, model, dpi int, eccLevel, input, charMode, text string) (err error) {
	if err = rec.HasError(ErrPort); err != nil {
		return
	}
	ret, _, _ := rec.sdk.Proc.QRCode.Call(rec.id,
		uintptr(xPos),
		uintptr(yPos),
		uintptr(orientation),
		uintptr(model),
		uintptr(dpi),
		String2UintPTR(eccLevel),
		String2UintPTR(input),
		String2UintPTR(charMode),
		String2UintPTR(text),
	)
	if ret != 0 {
		err = fmt.Errorf("join qrcode failed,err code:%v", int32(ret))
		return
	}
	return
}

func (rec *Printer) JoinDataMatrixBarcode(xPos, yPos, orientation, codeHeight, level, columns, rows, formatId, aspectRatio int, text string) (err error) {
	if err = rec.HasError(ErrPort); err != nil {
		return
	}
	ret, _, _ := rec.sdk.Proc.DataMatrixBarCode.Call(rec.id,
		uintptr(xPos),
		uintptr(yPos),
		uintptr(orientation),
		uintptr(codeHeight),
		uintptr(level),
		uintptr(columns),
		uintptr(rows),
		uintptr(formatId),
		uintptr(aspectRatio),
		String2UintPTR(text),
	)
	if ret != 0 {
		err = fmt.Errorf("join data matrix barcode failed,err code:%v", int32(ret))
		return
	}
	return
}

func (rec *Printer) JoinImage(xPos, yPos int, imgName string) (err error) {
	if err = rec.HasError(ErrPort); err != nil {
		return
	}
	ret, _, _ := rec.sdk.Proc.QRCode.Call(rec.id,
		uintptr(xPos),
		uintptr(yPos),
		String2UintPTR(imgName),
	)
	if ret != 0 {
		err = fmt.Errorf("join image failed,err code:%v", int32(ret))
		return
	}
	return
}

func (rec *Printer) JoinGraphicBox(xPos, yPos, width, height, thickness, rounding int) (err error) {
	if err = rec.HasError(ErrPort); err != nil {
		return
	}
	ret, _, _ := rec.sdk.Proc.GraphicBox.Call(rec.id,
		uintptr(xPos),
		uintptr(yPos),
		uintptr(width),
		uintptr(height),
		uintptr(thickness),
		uintptr(rounding),
	)
	if ret != 0 {
		err = fmt.Errorf("join graphic box failed,err code:%v", int32(ret))
		return
	}
	return
}
func (rec *Printer) PrintConfigurationLabel() (err error) {
	if err = rec.HasError(ErrPort); err != nil {
		return
	}
	ret, _, _ := rec.sdk.Proc.PrintConfigurationLabel.Call(rec.id)
	if ret != 0 {
		err = fmt.Errorf("print failed,err code:%v", int32(ret))
		return
	}
	return
}

func (rec *Printer) GetPrinterStatus() (value int, err error) {
	if err = rec.HasError(ErrPort); err != nil {
		return
	}
	var buffer uintptr
	ret, _, _ := rec.sdk.Proc.GetPrinterStatus.Call(rec.id, uintptr(unsafe.Pointer(&buffer)))
	if ret < 0 {
		err = fmt.Errorf("get printer status failed,err code:%v", int32(ret))
		return
	}
	return StatusDecode(buffer), nil
}
