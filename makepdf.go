package main

import (
    "github.com/signintech/gopdf"
    "fmt"
    "os"
    "io"
    "encoding/csv"
)

// 構造体定義
type Page struct {
    Title string
    Contents []string
}
type Presentation struct {
    Pages []Page
}

func main() {

    // 定数定義
    path := "sample.csv"
    presentation := new (Presentation)
    
    // CSVファイル読み込み
    fp, err := os.Open(path)
    if err != nil {
        panic(err)
    }
    defer fp.Close()

    reader := csv.NewReader(fp)
    // エラーになるまで繰り返し
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        } else if err != nil {
            panic(err)
        }

        //構造体変換処理
        page := Page{Title: record[0], Contents: record[1:]}
        fmt.Println(page)
        presentation.Pages = append(presentation.Pages, page)
    }

    // gopdf のオブジェクトを作成 --- (*1)
    pdf := gopdf.GoPdf{}
    // A4(横)のページを作る --- (*2)
    A4 := *gopdf.PageSizeA4
    A4Yoko := gopdf.Rect{W: A4.H, H: A4.W}
    pdf.Start(gopdf.Config{PageSize: A4Yoko})
    // TTFフォントを取り込む --- (*3)
    err = pdf.AddTTFFont("mukasi", "./gomarice_mukasi_mukasi.ttf"    )
    if err != nil {
        panic(err)
    }
    err = pdf.SetFont("mukasi", "", 26) // フォントサイズを選択
    if err != nil {
        panic(err)
    }
    // PDFへPage.Contentsが入るY軸開始範囲
    minY := 200
    // 初期座標
    y := 0

    //　文字を書き込む --- (*4)
    for _, v := range presentation.Pages {

        pdf.AddPage()
        pdf.SetX(80)
        pdf.SetY(100)
        pdf.Cell(nil, v.Title)
        y = minY
        // Contentsを書き込む
        for _, z := range v.Contents {
            pdf.SetX(80)
            pdf.SetY(float64(y))
            pdf.Cell(nil, z)
            y = y + 50
        }
    }
    // PDFをファイルに書き出す --- (*5)
    pdf.WritePdf("hello2.pdf")
}
