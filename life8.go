package main

import (
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "image/gif"
    "math/rand"
    "time"
    "math"
    "os"
    "strconv"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Changeable interface {
    Set(x,y int, c color.Color)
}




var (
  white color.Color = color.RGBA{255, 255, 255, 255}
  black color.Color = color.RGBA{0, 0, 0, 255}
  blue  color.Color = color.RGBA{0, 0, 255, 255}
)

func Round(f float64) float64 {
    return math.Floor(f + .5)
}

func collidecheck(checkmatrix [][]float64) (int, int, string) {

    collision := "not sure"
    returni := 0
    returnj := 0

    for i:=0; i<len(checkmatrix); i++{
        for j:=0; j<len(checkmatrix); j++{
            if i!=j && checkmatrix[i][0] == checkmatrix[j][0] && checkmatrix[i][1] == checkmatrix[j][1]{
                fmt.Println("collision struck! at i,j:", i, " ", j)
                returni = i
                returnj = j
                collision = "collision!"
                return returni, returnj, collision
            }
        }
    }

    collision = "no collision!"
    return returni, returnj, collision
}

func jitter(jittermatrix [][]float64, picturesize float64) [][]float64{

	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    colori := 0
    colorj := 0
    collisioncounter := 0

 	for intvariable:=0; intvariable<len(jittermatrix);{

        positionxnew := jittermatrix[intvariable][0]
        positionynew := jittermatrix[intvariable][1]
	
		for m:= 0; m<1;{
			    s1 = rand.NewSource(time.Now().UnixNano())
        	r1 = rand.New(s1)
            
          positionxnew = jittermatrix[intvariable][0]
          positionynew = jittermatrix[intvariable][1]

        	positionxnew = (positionxnew + float64(Round((Round(r1.Float64())*-2+1))))
        	positionynew = (positionynew + float64(Round((Round(r1.Float64())*-2+1))))


        	if (positionxnew > 0 && positionynew < picturesize) && (positionynew > 0 && positionynew < picturesize) {
        	   m = 1
        	   jittermatrix[intvariable][0] = positionxnew
        	   jittermatrix[intvariable][1] = positionynew 

        	   doesitcollide := ""

        	   colori,colorj,doesitcollide=collidecheck(jittermatrix)
        	   if doesitcollide=="collision!"{
        	       collisioncounter = collisioncounter + 1
                   jittermatrix[colori][2]+=1
                   jittermatrix[colori][3]+=1
                   jittermatrix[colorj][2]+=1
                   jittermatrix[colorj][3]+=1
        	   } else{
        	       collisioncounter = 0
                   intvariable = intvariable + 1
        	   }

        	   if collisioncounter >= 5{
        		   jittermatrix = readjustsize(len(jittermatrix)-1, jittermatrix)
        	   }
    	   }
        }
   	}

   	return jittermatrix
}


func readjustsize (maxsize int, matrixtoadjust [][]float64)[][]float64{

	dummyindex1 := len(matrixtoadjust) - 1
	dummyindex2 := len(matrixtoadjust)

	for len(matrixtoadjust) > maxsize{
		matrixtoadjust = append(matrixtoadjust[:dummyindex1], matrixtoadjust[dummyindex2:]...)
	}

	return matrixtoadjust
}

func addrow (rowtoadd [][]float64, matrixtoaddto [][]float64) [][]float64 {

	matrixtoaddto = append(matrixtoaddto, rowtoadd...)

	return matrixtoaddto

}

func converttoslice (firstindex float64, secondindex float64, thirdindex float64, fourthindex float64) []float64 {

	returnslice := make([]float64, 0)
	returnslice = append(returnslice, firstindex)
	returnslice = append(returnslice, secondindex)
	returnslice = append(returnslice, thirdindex)
	returnslice = append(returnslice, fourthindex)

	return returnslice

}


func drawmatrixgif (drawmatrix [][]float64, matrixsize float64, picnum int) {

  //outGif := &gif.GIF{}

  if picnum == 0 {
    
      img := image.NewRGBA(image.Rect(0,0,int(matrixsize), int(matrixsize)))
      draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.ZP, draw.Src) 

      for i := 0; i < int(matrixsize); i++ {
          img.Set(int(drawmatrix[i][0]), int(drawmatrix[i][1]), color.RGBA{100+uint8(drawmatrix[i][2])+uint8(drawmatrix[i][3]),100+uint8(drawmatrix[i][2]),100+uint8(drawmatrix[i][3]),255})           
      }

 //     func main() {
 // f, err := os.OpenFile("foo.gif", os.O_CREATE|os.O_RDWR, 0x666)
 // if err != nil {
 //   log.Fatal(err)
 // }
 // r := image.Rect(0, 0, 240, 320)
 // i := image.NewRGBA(r)
 // err = gif.Encode(f, i, nil)
 // if err != nil {
 //   log.Fatal(err)
 // }
 //}

      if imgFileGif, err := os.OpenFile("out0.gif", os.O_CREATE|os.O_RDWR, 0x666); err != nil {
         fmt.Println("Gif error: ", err)
      } else {
         defer imgFileGif.Close()
         gif.Encode(imgFileGif, img, nil)
      }

  }else{

      picnum = picnum-1
      s:= strconv.Itoa(picnum)
      var openfilename string = "out" + s + ".gif"

      fmt.Println(openfilename)

      imgFileGif, err := os.OpenFile(openfilename, os.O_CREATE|os.O_RDWR,0x666)

      if err != nil {
          fmt.Println(err, "1")
          fmt.Println("error 1 heyo")
      }


      img := image.NewRGBA(image.Rect(0,0,int(matrixsize), int(matrixsize)))
      draw.Draw(img, img.Bounds(), &image.Uniform{white}, image.ZP, draw.Src) 

      for i := 0; i < int(matrixsize); i++ {
          img.Set(int(drawmatrix[i][0]), int(drawmatrix[i][1]), color.RGBA{100+uint8(drawmatrix[i][2])+uint8(drawmatrix[i][3]),100+uint8(drawmatrix[i][2]),100+uint8(drawmatrix[i][3]),255})           
      }

      s = strconv.Itoa(picnum)

      f, _ := os.OpenFile("out"+s+".gif", os.O_RDWR|os.O_CREATE, 0x600)
      defer f.Close()
      gif.Encode(imgFileGif, img, nil)


     // img1, err := gif.Decode(imgFile1)
     // if err != nil {
     //     fmt.Println(err, "2")
     //     fmt.Println("error 2 heyo")
     // }


     // for i:= 0; i < int(matrixsize); i++{
     //     if cimg, ok := img1.(Changeable); ok {
     //         for i := 0; i < int(matrixsize); i++ {
     //             cimg.Set(int(drawmatrix[i][0]), int(drawmatrix[i][1]), color.RGBA{255,uint8(drawmatrix[i][2]),uint8(drawmatrix[i][3]),255})           
     //         }
     //     }
     // }

  }
  

}




func makematrix(matrixrows int, picturesize float64) [][]float64{
	    dotmatrix := make([][]float64, matrixrows)
    for i := 0; i < matrixrows; i++ {
        dotmatrix[i] = make([]float64, 4)
    }

    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)    

    for m := 0; m < matrixrows; m++{
  
        dotmatrix[m][0] = math.Abs(Round(picturesize*r1.Float64()))
        
        time.Sleep(1 * time.Millisecond)
        s1 = rand.NewSource(time.Now().UnixNano())
    	  r1 = rand.New(s1)
        
        dotmatrix[m][0] = math.Abs(dotmatrix[m][0]-Round(picturesize/2*r1.Float64()))
  		
  	  	time.Sleep(1 * time.Millisecond)
        s1 = rand.NewSource(time.Now().UnixNano())
      	r1 = rand.New(s1)

  		  dotmatrix[m][1] = math.Abs(Round(picturesize*r1.Float64()))
       	
       	time.Sleep(1 * time.Millisecond)
        s1 = rand.NewSource(time.Now().UnixNano())
    	  r1 = rand.New(s1)

        dotmatrix[m][1] = math.Abs(dotmatrix[m][1]-Round(picturesize/2*r1.Float64()))
  		
        dotmatrix[m][2] = 0
        dotmatrix[m][3] = 0

        time.Sleep(25 * time.Millisecond)

    }

    return dotmatrix
}


func main(){


  //variable declarations
  dotmatrixsize:=1000
  dotmatrix := make([][]float64,dotmatrixsize)
  var looptimes int = 255
  var picturesize float64 = 1200

  dotmatrix = makematrix(dotmatrixsize,picturesize)

  fmt.Println(dotmatrix)

  //main loop

  for i:= 0; i<looptimes; i++{
    dotmatrix = jitter(dotmatrix, picturesize)
    drawmatrixgif(dotmatrix,float64(len(dotmatrix)),i)
    fmt.Println(dotmatrix)
  }


  //here im stitching together all of the static gif images



  files := make([]string, 0)

  for as:= 0; as<looptimes; as++{

      fmt.Println(as)
      s:= strconv.Itoa(as)
      fileappend := make([]string, 1)
      fmt.Println(fileappend)
      fileappend[0] = "out" + s + ".gif"
      files = append(files, fileappend...)
  }

    fmt.Println("hey there good looking")
    fmt.Println(files)

    // load static image and construct outGif
    outGif := &gif.GIF{}
    for _, name := range files {
        f, _ := os.Open(name)
        inGif, _ := gif.Decode(f)
        f.Close()

        outGif.Image = append(outGif.Image, inGif.(*image.Paletted))
        outGif.Delay = append(outGif.Delay, 0)
    }

    // save to out.gif
    f, _ := os.OpenFile("finalout.gif", os.O_WRONLY|os.O_CREATE, 0600)
    defer f.Close()
    gif.EncodeAll(f, outGif)








}