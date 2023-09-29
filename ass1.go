
//THIS IS LUCKY TOM 1/1000 luck
//              ..,,:::;;iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii;::1ffLLLLLLLLLLL
//             ..,,::;iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii;::::iiiiiiii;ii11ttttttttt
//          .,,,:;;;;iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii:. ,:;iiiiiiiiiiiii;;;;;;;;;;
//        .,..,:;iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii:. ,t1iiiiiiiiiiiiiiiiiiiiiiiii
//           :iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii;,  ;C00L1iiiiiiiiiiiiiiiiiiiiiii
//          :iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii;,  :tG0000Gfiiiiiiiiiiiiiiiiiiiiii
//::::::,,..;iiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii:. ,tG00000000Liiiiiiiiiiiiiiiiiiiii
//;;;;iiii::iiiiiiiiiiiiiiiiiiiii;iiiiiiiiii;.  :C00000000000C1iiiiiiiiiiiiiiiiiii
//11i;::;i;;iiiiiiiiiiiiiiii;;;iiii;iiiiii;,  :f0Ct11fC0000000G1;iiiiiiii;;;;i:;ii
//11111i;:;iiiiiiiiiiiiiiiiiiiiiiii:;iii;,  ,tG00i,. .,iL000000L,:iii;;;;itlCf:ii;
//1111111i:;iiiiiiiiiiiiiiiiiiii;;iii;:,..,:t0000t,    .:l0000C;,,;;;ifCG88@@L:;;:
//11111111i;iiiiiiiiiiiiiiiiiiiiiii;;;:::;ii;f000Gt:.. .:f000L;,::;tG8@@@@88Gt:;;i
//11111111i;iiiiiiiiiiiiiiiiiiiiiii,.:,::;;ii;f0000Cft1tLG0Gf::,,iC8@@8@80CC0C;iii
//11111111;;iiiiiiiiiiiiiiiiiiii;,.  :;;:;:;ii;1C00000000Gli::,iC8@88@80CG8@@81;;i
//1111iii;;iiiiiiiiiiiiiiiiiii;,  ..:iiii;:;iii;itCGGGCLti;:::f8@888@8CC8@@@@@8L1i
//1111;;;;;;;;iiiiiiiiiiiiii;.  .iLCt;;ii;;iiiiii;;i;:::,:::;L8@888@0L0@@@@@@@@8Gf
//1111111i;;:;iiiiiiiiiiii;,  .1C0000Gt;;iiiiiiiiiii;:;;;;;;G@8888@GC8@@@@@@@8GG08
//111111i;;;;:iiiiiiiiiii:.  :LGLLLCG00C1;;iiiiiiiiiiiiii;1G@888@@CC@@@@@@@0CG8@@@
//i1111111111;;iiiiiiiii:  .1GL;,..,iL000Li;iiiiiiiiiiii;i0@888@@0C@@@@@@@CC8@@@@@
//.,;i11111111;;iiiiiii:  ;L001,    .:f000Gi:iiiiiiiii;;;C@8888@@@@@@@@8GC8@@@@@@@
//   .,:;iiiii;:;iiiii:  iGG00L:.    .;G000L,;iiii;::::::C8@@8@@@@@@@@0C0@@@@@@@@@
//        ..... .:iiii: .tGGG00L1;,..,iGGGGf.:;iii;,..,. .:f88@@@@@@@CC8@@@@@@@@@@
//                :iii;.:;tGGGG00GCLffCGGGL:,:iiii;..f0i    L@@@@@@@@8@@@@@88@@@@@
//                 ,:ii;ii;1LGGGGG00GGGGGf;::;ii;;, ;8t.    f@@@@@@@@@@@@@@@8@@@@@
//                   ,;iiiii;1LGGGGGGGGLi,:;::i1tfi .:      C@@@@@@@@@8GLC8@@@@@@@
//                    .:iiiii;i1fLCCCCt:,:;;:1C08@0i:::::;;i0@@@@@@@0LiiG@@@@@@@@@
//                      ,;iiiiii;iii;:,:;;;;t8@@8@@@88888@@@@@8888G1::f8@@@@@@@@@@
//                       .:iiiiiiii;;;;iii;f8@@@@@@@@@@@@@@@@@@8Gi,,t08G0@@@@@@@@@
//                         ,;iiiiiiiiiiii;i8@80GG@@@@@@@@@@@@@@8t:1G80G88888@@@@@@
//                          .:iiiiiiiiiii;180CC0@@@@@@@@@@@@@@@CL08008@@888@@@@@@@
//                            :iiiiiiiiii;;CG88@@@@@8GG8@@@@@@@@@@@@@888@@@@@@@@@@
//                            ,iiiiiiii;:::0@@8@@8GGG0@@@@@@888@@@@@@@@@@@@@@@@@@@

package main

import (
	"fmt"
)



type WeatherProvider interface {
	Fetch() (float64, float64)
}

type OpenWeatherMap struct {
	ApiKey string
}

func (owm *OpenWeatherMap) Fetch() (float64, float64) {
	
	return 22.5, 45.0 
}

type DarkSky struct {
	ApiKey string
}

func (ds DarkSky) Fetch() (float64, float64) {
	
	return 20.8, 55.2 
}



type Observer interface {
	Update(float64, float64)
}

type WeatherStation struct {
	Observers []Observer
	Provider  WeatherProvider
}

func (ws *WeatherStation) AddObserver(o Observer) {
	ws.Observers = append(ws.Observers, o)
}

func (ws *WeatherStation) RemoveObserver(o Observer) {
	for i, observer := range ws.Observers {
		if observer == o {
			ws.Observers = append(ws.Observers[:i], ws.Observers[i+1:]...)
			break
		}
	}
}

func (ws *WeatherStation) NotifyObservers() {
	temperature, humidity := ws.Provider.Fetch()
	for _, observer := range ws.Observers {
		observer.Update(temperature, humidity)
	}
}



type Display struct {
	Name string
}

func (d *Display) Update(temperature, humidity float64) {
	fmt.Printf("%s: Temperature %.2f°C, Humidity %.2f%%\n", d.Name, temperature, humidity)
}

func main() {
	
	openWeatherMap := &OpenWeatherMap{ApiKey: "your_openweathermap_api_key"}
	darkSky := DarkSky{ApiKey: "your_darksky_api_key"}

	
	weatherStation := &WeatherStation{Provider: openWeatherMap}


	display1 := &Display{Name: "Display 1"}
	display2 := &Display{Name: "Display 2"}

	
	weatherStation.AddObserver(display1)
	weatherStation.AddObserver(display2)

	
	fmt.Println("Fetching weather data from OpenWeatherMap:")
	weatherStation.NotifyObservers()

	weatherStation.Provider = darkSky
	fmt.Println("\nFetching weather data from DarkSky:")
	weatherStation.NotifyObservers()
}
