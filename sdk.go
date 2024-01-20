package zpl_printer

import (
	"errors"
	"fmt"

	"golang.org/x/sys/windows"
)

type Proc struct {
	InitPrinter             *windows.Proc // implemented
	ReleasePrinter          *windows.Proc // implemented
	GetPrinterStatus        *windows.Proc // implemented
	OpenPort                *windows.Proc // implemented
	ClosePort               *windows.Proc // implemented
	StartFormat             *windows.Proc // implemented
	EndFormat               *windows.Proc // implemented
	Text                    *windows.Proc // implemented
	TextBlock               *windows.Proc // implemented
	BarCode128              *windows.Proc // implemented
	DataMatrixBarCode       *windows.Proc // implemented
	QRCode                  *windows.Proc // implemented
	GraphicBox              *windows.Proc // implemented
	PrintImage              *windows.Proc // implemented
	PrintConfigurationLabel *windows.Proc // implemented
}
type Sdk struct {
	dllSdk *windows.DLL
	Proc   *Proc
}

func NewSDK(dllPath string) (sdk *Sdk, err error) {
	sdk = &Sdk{}
	sdk.dllSdk, err = windows.LoadDLL(dllPath)
	if err != nil {
		err = fmt.Errorf("[InitDLL] Load dll file %v failed,err:%v", dllPath, err.Error())
	}
	p := &Proc{}
	var procError []error // 切片记录加载错误
	p.InitPrinter, err = sdk.dllSdk.FindProc("InitPrinter")
	if err != nil {
		procError = append(procError, err)
	}
	p.ReleasePrinter, err = sdk.dllSdk.FindProc("ReleasePrinter")
	if err != nil {
		procError = append(procError, err)
	}
	p.GetPrinterStatus, err = sdk.dllSdk.FindProc("ZPL_GetPrinterStatus")
	if err != nil {
		procError = append(procError, err)
	}
	p.OpenPort, err = sdk.dllSdk.FindProc("OpenPort")
	if err != nil {
		procError = append(procError, err)
	}
	p.ClosePort, err = sdk.dllSdk.FindProc("ClosePort")
	if err != nil {
		procError = append(procError, err)
	}
	p.StartFormat, err = sdk.dllSdk.FindProc("ZPL_StartFormat")
	if err != nil {
		procError = append(procError, err)
	}
	p.EndFormat, err = sdk.dllSdk.FindProc("ZPL_EndFormat")
	if err != nil {
		procError = append(procError, err)
	}
	p.Text, err = sdk.dllSdk.FindProc("ZPL_Text")
	if err != nil {
		procError = append(procError, err)
	}
	p.TextBlock, err = sdk.dllSdk.FindProc("ZPL_Text_Block")
	if err != nil {
		procError = append(procError, err)
	}
	p.BarCode128, err = sdk.dllSdk.FindProc("ZPL_BarCode128")
	if err != nil {
		procError = append(procError, err)
	}
	p.QRCode, err = sdk.dllSdk.FindProc("ZPL_QRCode")
	if err != nil {
		procError = append(procError, err)
	}
	p.GraphicBox, err = sdk.dllSdk.FindProc("ZPL_GraphicBox")
	if err != nil {
		procError = append(procError, err)
	}
	p.PrintImage, err = sdk.dllSdk.FindProc("ZPL_PrintImage")
	if err != nil {
		procError = append(procError, err)
	}
	p.DataMatrixBarCode, err = sdk.dllSdk.FindProc("ZPL_DataMatrixBarcode")
	if err != nil {
		procError = append(procError, err)
	}
	p.GetPrinterStatus, err = sdk.dllSdk.FindProc("ZPL_GetPrinterStatus")
	if err != nil {
		procError = append(procError, err)
	}
	p.PrintConfigurationLabel, err = sdk.dllSdk.FindProc("ZPL_PrintConfigurationLabel")
	if err != nil {
		procError = append(procError, err)
	}
	if len(procError) != 0 {
		fmt.Println(procError)
		err = errors.New("[InitDLL] Find proc error")
		_ = sdk.Release()
		return nil, err
	}
	sdk.Proc = p
	return
}

func (rec *Sdk) Release() (err error) {
	if rec.dllSdk == nil {
		return
	}
	err = rec.dllSdk.Release()
	if err != nil {
		err = fmt.Errorf("release proc failed,err : %v", err.Error())
		fmt.Println(err)
		return err
	}
	rec.Proc = nil
	rec.dllSdk = nil
	return
}
