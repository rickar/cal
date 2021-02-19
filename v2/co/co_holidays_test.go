// (c) Rick Arnold. Licensed under the BSD license (see LICENSE).
// (c) Juan P. Garcia

package co

import (
	"testing"
	"time"

	"github.com/rickar/cal/v2"
)

func d(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, cal.DefaultLoc)
}

func TestHolidays(t *testing.T) {
	tests := []struct {
		h       *cal.Holiday
		y       int
		wantAct time.Time
		wantObs time.Time
	}{
		{AñoNuevo, 2015, d(2015, 1, 1), d(2015, 1, 1)},
		{AñoNuevo, 2016, d(2016, 1, 1), d(2016, 1, 1)},
		{AñoNuevo, 2017, d(2017, 1, 1), d(2017, 1, 1)},
		{AñoNuevo, 2018, d(2018, 1, 1), d(2018, 1, 1)},
		{AñoNuevo, 2019, d(2019, 1, 1), d(2019, 1, 1)},
		{AñoNuevo, 2020, d(2020, 1, 1), d(2020, 1, 1)},
		{AñoNuevo, 2021, d(2021, 1, 1), d(2021, 1, 1)},
		{AñoNuevo, 2022, d(2022, 1, 1), d(2022, 1, 1)},

		{Reyes, 2015, d(2015, 1, 12), d(2015, 1, 12)},
		{Reyes, 2016, d(2016, 1, 11), d(2016, 1, 11)},
		{Reyes, 2017, d(2017, 1, 9), d(2017, 1, 9)},
		{Reyes, 2018, d(2018, 1, 8), d(2018, 1, 8)},
		{Reyes, 2019, d(2019, 1, 7), d(2019, 1, 7)},
		{Reyes, 2020, d(2020, 1, 6), d(2020, 1, 6)},
		{Reyes, 2021, d(2021, 1, 11), d(2021, 1, 11)},
		{Reyes, 2022, d(2022, 1, 10), d(2022, 1, 10)},

		{DomingoDeRamos, 2015, d(2015, 3, 29), d(2015, 3, 29)},
		{DomingoDeRamos, 2016, d(2016, 3, 20), d(2016, 3, 20)},
		{DomingoDeRamos, 2017, d(2017, 4, 9), d(2017, 4, 9)},
		{DomingoDeRamos, 2018, d(2018, 3, 25), d(2018, 3, 25)},
		{DomingoDeRamos, 2019, d(2019, 4, 14), d(2019, 4, 14)},
		{DomingoDeRamos, 2020, d(2020, 4, 5), d(2020, 4, 5)},
		{DomingoDeRamos, 2021, d(2021, 3, 28), d(2021, 3, 28)},
		{DomingoDeRamos, 2022, d(2022, 4, 10), d(2022, 4, 10)},

		{JuevesSanto, 2015, d(2015, 4, 2), d(2015, 4, 2)},
		{JuevesSanto, 2016, d(2016, 3, 24), d(2016, 3, 24)},
		{JuevesSanto, 2017, d(2017, 4, 13), d(2017, 4, 13)},
		{JuevesSanto, 2018, d(2018, 3, 29), d(2018, 3, 29)},
		{JuevesSanto, 2019, d(2019, 4, 18), d(2019, 4, 18)},
		{JuevesSanto, 2020, d(2020, 4, 9), d(2020, 4, 9)},
		{JuevesSanto, 2021, d(2021, 4, 1), d(2021, 4, 1)},
		{JuevesSanto, 2022, d(2022, 4, 14), d(2022, 4, 14)},

		{ViernesSanto, 2015, d(2015, 4, 3), d(2015, 4, 3)},
		{ViernesSanto, 2016, d(2016, 3, 25), d(2016, 3, 25)},
		{ViernesSanto, 2017, d(2017, 4, 14), d(2017, 4, 14)},
		{ViernesSanto, 2018, d(2018, 3, 30), d(2018, 3, 30)},
		{ViernesSanto, 2019, d(2019, 4, 19), d(2019, 4, 19)},
		{ViernesSanto, 2020, d(2020, 4, 10), d(2020, 4, 10)},
		{ViernesSanto, 2021, d(2021, 4, 2), d(2021, 4, 2)},
		{ViernesSanto, 2022, d(2022, 4, 15), d(2022, 4, 15)},

		{Pascua, 2015, d(2015, 4, 5), d(2015, 4, 5)},
		{Pascua, 2016, d(2016, 3, 27), d(2016, 3, 27)},
		{Pascua, 2017, d(2017, 4, 16), d(2017, 4, 16)},
		{Pascua, 2018, d(2018, 4, 1), d(2018, 4, 1)},
		{Pascua, 2019, d(2019, 4, 21), d(2019, 4, 21)},
		{Pascua, 2020, d(2020, 4, 12), d(2020, 4, 12)},
		{Pascua, 2021, d(2021, 4, 4), d(2021, 4, 4)},
		{Pascua, 2022, d(2022, 4, 17), d(2022, 4, 17)},

		{DíaAscension, 2015, d(2015, 5, 18), d(2015, 5, 18)},
		{DíaAscension, 2016, d(2016, 5, 9), d(2016, 5, 9)},
		{DíaAscension, 2017, d(2017, 5, 29), d(2017, 5, 29)},
		{DíaAscension, 2018, d(2018, 5, 14), d(2018, 5, 14)},
		{DíaAscension, 2019, d(2019, 6, 3), d(2019, 6, 3)},
		{DíaAscension, 2020, d(2020, 5, 25), d(2020, 5, 25)},
		{DíaAscension, 2021, d(2021, 5, 17), d(2021, 5, 17)},
		{DíaAscension, 2022, d(2022, 5, 30), d(2022, 5, 30)},

		{CorpusChristi, 2015, d(2015, 6, 8), d(2015, 6, 8)},
		{CorpusChristi, 2016, d(2016, 5, 30), d(2016, 5, 30)},
		{CorpusChristi, 2017, d(2017, 6, 19), d(2017, 6, 19)},
		{CorpusChristi, 2018, d(2018, 6, 4), d(2018, 6, 4)},
		{CorpusChristi, 2019, d(2019, 6, 24), d(2019, 6, 24)},
		{CorpusChristi, 2020, d(2020, 6, 15), d(2020, 6, 15)},
		{CorpusChristi, 2021, d(2021, 6, 7), d(2021, 6, 7)},
		{CorpusChristi, 2022, d(2022, 6, 20), d(2022, 6, 20)},

		{SagradoCorazon, 2015, d(2015, 6, 15), d(2015, 6, 15)},
		{SagradoCorazon, 2016, d(2016, 6, 6), d(2016, 6, 6)},
		{SagradoCorazon, 2017, d(2017, 6, 26), d(2017, 6, 26)},
		{SagradoCorazon, 2018, d(2018, 6, 11), d(2018, 6, 11)},
		{SagradoCorazon, 2019, d(2019, 7, 1), d(2019, 7, 1)},
		{SagradoCorazon, 2020, d(2020, 6, 22), d(2020, 6, 22)},
		{SagradoCorazon, 2021, d(2021, 6, 14), d(2021, 6, 14)},
		{SagradoCorazon, 2022, d(2022, 6, 27), d(2022, 6, 27)},

		{DíaMujer, 2015, d(2015, 3, 8), d(2015, 3, 8)},
		{DíaMujer, 2016, d(2016, 3, 8), d(2016, 3, 8)},
		{DíaMujer, 2017, d(2017, 3, 8), d(2017, 3, 8)},
		{DíaMujer, 2018, d(2018, 3, 8), d(2018, 3, 8)},
		{DíaMujer, 2019, d(2019, 3, 8), d(2019, 3, 8)},
		{DíaMujer, 2020, d(2020, 3, 8), d(2020, 3, 8)},
		{DíaMujer, 2021, d(2021, 3, 8), d(2021, 3, 8)},
		{DíaMujer, 2022, d(2022, 3, 8), d(2022, 3, 8)},

		{DíaSanJose, 2015, d(2015, 3, 23), d(2015, 3, 23)},
		{DíaSanJose, 2016, d(2016, 3, 21), d(2016, 3, 21)},
		{DíaSanJose, 2017, d(2017, 3, 20), d(2017, 3, 20)},
		{DíaSanJose, 2018, d(2018, 3, 19), d(2018, 3, 19)},
		{DíaSanJose, 2019, d(2019, 3, 25), d(2019, 3, 25)},
		{DíaSanJose, 2020, d(2020, 3, 23), d(2020, 3, 23)},
		{DíaSanJose, 2021, d(2021, 3, 22), d(2021, 3, 22)},
		{DíaSanJose, 2022, d(2022, 3, 21), d(2022, 3, 21)},

		{DíaIdioma, 2015, d(2015, 4, 23), d(2015, 4, 23)},
		{DíaIdioma, 2016, d(2016, 4, 23), d(2016, 4, 23)},
		{DíaIdioma, 2017, d(2017, 4, 23), d(2017, 4, 23)},
		{DíaIdioma, 2018, d(2018, 4, 23), d(2018, 4, 23)},
		{DíaIdioma, 2019, d(2019, 4, 23), d(2019, 4, 23)},
		{DíaIdioma, 2020, d(2020, 4, 23), d(2020, 4, 23)},
		{DíaIdioma, 2021, d(2021, 4, 23), d(2021, 4, 23)},
		{DíaIdioma, 2022, d(2022, 4, 23), d(2022, 4, 23)},

		{DíaNino, 2015, d(2015, 4, 25), d(2015, 4, 25)},
		{DíaNino, 2016, d(2016, 4, 30), d(2016, 4, 30)},
		{DíaNino, 2017, d(2017, 4, 29), d(2017, 4, 29)},
		{DíaNino, 2018, d(2018, 4, 28), d(2018, 4, 28)},
		{DíaNino, 2019, d(2019, 4, 27), d(2019, 4, 27)},
		{DíaNino, 2020, d(2020, 4, 25), d(2020, 4, 25)},
		{DíaNino, 2021, d(2021, 4, 24), d(2021, 4, 24)},
		{DíaNino, 2022, d(2022, 4, 30), d(2022, 4, 30)},

		{Trabajo, 2015, d(2015, 5, 1), d(2015, 5, 1)},
		{Trabajo, 2016, d(2016, 5, 1), d(2016, 5, 1)},
		{Trabajo, 2017, d(2017, 5, 1), d(2017, 5, 1)},
		{Trabajo, 2018, d(2018, 5, 1), d(2018, 5, 1)},
		{Trabajo, 2019, d(2019, 5, 1), d(2019, 5, 1)},
		{Trabajo, 2020, d(2020, 5, 1), d(2020, 5, 1)},
		{Trabajo, 2021, d(2021, 5, 1), d(2021, 5, 1)},
		{Trabajo, 2022, d(2022, 5, 1), d(2022, 5, 1)},

		{DíaMadre, 2015, d(2015, 5, 10), d(2015, 5, 10)},
		{DíaMadre, 2016, d(2016, 5, 8), d(2016, 5, 8)},
		{DíaMadre, 2017, d(2017, 5, 14), d(2017, 5, 14)},
		{DíaMadre, 2018, d(2018, 5, 13), d(2018, 5, 13)},
		{DíaMadre, 2019, d(2019, 5, 12), d(2019, 5, 12)},
		{DíaMadre, 2020, d(2020, 5, 10), d(2020, 5, 10)},
		{DíaMadre, 2021, d(2021, 5, 9), d(2021, 5, 9)},
		{DíaMadre, 2022, d(2022, 5, 8), d(2022, 5, 8)},

		{DíaPadre, 2015, d(2015, 6, 21), d(2015, 6, 21)},
		{DíaPadre, 2016, d(2016, 6, 19), d(2016, 6, 19)},
		{DíaPadre, 2017, d(2017, 6, 18), d(2017, 6, 18)},
		{DíaPadre2018, 2018, d(2018, 6, 24), d(2018, 6, 24)},
		{DíaPadre, 2019, d(2019, 6, 16), d(2019, 6, 16)},
		{DíaPadre, 2020, d(2020, 6, 21), d(2020, 6, 21)},
		{DíaPadre, 2021, d(2021, 6, 20), d(2021, 6, 20)},
		{DíaPadre, 2022, d(2022, 6, 19), d(2022, 6, 19)},

		{SanPedroSanPablo, 2015, d(2015, 6, 29), d(2015, 6, 29)},
		{SanPedroSanPablo, 2016, d(2016, 7, 4), d(2016, 7, 4)},
		{SanPedroSanPablo, 2017, d(2017, 7, 3), d(2017, 7, 3)},
		{SanPedroSanPablo, 2018, d(2018, 7, 2), d(2018, 7, 2)},
		{SanPedroSanPablo, 2019, d(2019, 7, 1), d(2019, 7, 1)},
		{SanPedroSanPablo, 2020, d(2020, 6, 29), d(2020, 6, 29)},
		{SanPedroSanPablo, 2021, d(2021, 7, 5), d(2021, 7, 5)},
		{SanPedroSanPablo, 2022, d(2022, 7, 4), d(2022, 7, 4)},

		{Independencia, 2015, d(2015, 7, 20), d(2015, 7, 20)},
		{Independencia, 2016, d(2016, 7, 20), d(2016, 7, 20)},
		{Independencia, 2017, d(2017, 7, 20), d(2017, 7, 20)},
		{Independencia, 2018, d(2018, 7, 20), d(2018, 7, 20)},
		{Independencia, 2019, d(2019, 7, 20), d(2019, 7, 20)},
		{Independencia, 2020, d(2020, 7, 20), d(2020, 7, 20)},
		{Independencia, 2021, d(2021, 7, 20), d(2021, 7, 20)},
		{Independencia, 2022, d(2022, 7, 20), d(2022, 7, 20)},

		{BatallaBoyaca, 2015, d(2015, 8, 7), d(2015, 8, 7)},
		{BatallaBoyaca, 2016, d(2016, 8, 7), d(2016, 8, 7)},
		{BatallaBoyaca, 2017, d(2017, 8, 7), d(2017, 8, 7)},
		{BatallaBoyaca, 2018, d(2018, 8, 7), d(2018, 8, 7)},
		{BatallaBoyaca, 2019, d(2019, 8, 7), d(2019, 8, 7)},
		{BatallaBoyaca, 2020, d(2020, 8, 7), d(2020, 8, 7)},
		{BatallaBoyaca, 2021, d(2021, 8, 7), d(2021, 8, 7)},
		{BatallaBoyaca, 2022, d(2022, 8, 7), d(2022, 8, 7)},

		{Asunción, 2015, d(2015, 8, 17), d(2015, 8, 17)},
		{Asunción, 2016, d(2016, 8, 15), d(2016, 8, 15)},
		{Asunción, 2017, d(2017, 8, 21), d(2017, 8, 21)},
		{Asunción, 2018, d(2018, 8, 20), d(2018, 8, 20)},
		{Asunción, 2019, d(2019, 8, 19), d(2019, 8, 19)},
		{Asunción, 2020, d(2020, 8, 17), d(2020, 8, 17)},
		{Asunción, 2021, d(2021, 8, 16), d(2021, 8, 16)},
		{Asunción, 2022, d(2022, 8, 15), d(2022, 8, 15)},

		{DíaAmorAmistad, 2015, d(2015, 9, 19), d(2015, 9, 19)},
		{DíaAmorAmistad, 2016, d(2016, 9, 17), d(2016, 9, 17)},
		{DíaAmorAmistad, 2017, d(2017, 9, 16), d(2017, 9, 16)},
		{DíaAmorAmistad, 2018, d(2018, 9, 15), d(2018, 9, 15)},
		{DíaAmorAmistad, 2019, d(2019, 9, 21), d(2019, 9, 21)},
		{DíaAmorAmistad, 2020, d(2020, 9, 19), d(2020, 9, 19)},
		{DíaAmorAmistad, 2021, d(2021, 9, 18), d(2021, 9, 18)},
		{DíaAmorAmistad, 2022, d(2022, 9, 17), d(2022, 9, 17)},

		{DíaRaza, 2015, d(2015, 10, 12), d(2015, 10, 12)},
		{DíaRaza, 2016, d(2016, 10, 17), d(2016, 10, 17)},
		{DíaRaza, 2017, d(2017, 10, 16), d(2017, 10, 16)},
		{DíaRaza, 2018, d(2018, 10, 15), d(2018, 10, 15)},
		{DíaRaza, 2019, d(2019, 10, 14), d(2019, 10, 14)},
		{DíaRaza, 2020, d(2020, 10, 12), d(2020, 10, 12)},
		{DíaRaza, 2021, d(2021, 10, 18), d(2021, 10, 18)},
		{DíaRaza, 2022, d(2022, 10, 17), d(2022, 10, 17)},

		{TodosLosSantos, 2015, d(2015, 11, 2), d(2015, 11, 2)},
		{TodosLosSantos, 2016, d(2016, 11, 7), d(2016, 11, 7)},
		{TodosLosSantos, 2017, d(2017, 11, 6), d(2017, 11, 6)},
		{TodosLosSantos, 2018, d(2018, 11, 5), d(2018, 11, 5)},
		{TodosLosSantos, 2019, d(2019, 11, 4), d(2019, 11, 4)},
		{TodosLosSantos, 2020, d(2020, 11, 2), d(2020, 11, 2)},
		{TodosLosSantos, 2021, d(2021, 11, 1), d(2021, 11, 1)},
		{TodosLosSantos, 2022, d(2022, 11, 7), d(2022, 11, 7)},

		{IndependenciaCartagena, 2015, d(2015, 11, 16), d(2015, 11, 16)},
		{IndependenciaCartagena, 2016, d(2016, 11, 14), d(2016, 11, 14)},
		{IndependenciaCartagena, 2017, d(2017, 11, 13), d(2017, 11, 13)},
		{IndependenciaCartagena, 2018, d(2018, 11, 12), d(2018, 11, 12)},
		{IndependenciaCartagena, 2019, d(2019, 11, 11), d(2019, 11, 11)},
		{IndependenciaCartagena, 2020, d(2020, 11, 16), d(2020, 11, 16)},
		{IndependenciaCartagena, 2021, d(2021, 11, 15), d(2021, 11, 15)},
		{IndependenciaCartagena, 2022, d(2022, 11, 14), d(2022, 11, 14)},

		{DíaMujerColombiana, 2015, d(2015, 11, 14), d(2015, 11, 14)},
		{DíaMujerColombiana, 2016, d(2016, 11, 14), d(2016, 11, 14)},
		{DíaMujerColombiana, 2017, d(2017, 11, 14), d(2017, 11, 14)},
		{DíaMujerColombiana, 2018, d(2018, 11, 14), d(2018, 11, 14)},
		{DíaMujerColombiana, 2019, d(2019, 11, 14), d(2019, 11, 14)},
		{DíaMujerColombiana, 2020, d(2020, 11, 14), d(2020, 11, 14)},
		{DíaMujerColombiana, 2021, d(2021, 11, 14), d(2021, 11, 14)},
		{DíaMujerColombiana, 2022, d(2022, 11, 14), d(2022, 11, 14)},

		{VísperaInmaculadaConcepción, 2015, d(2015, 12, 7), d(2015, 12, 7)},
		{VísperaInmaculadaConcepción, 2016, d(2016, 12, 7), d(2016, 12, 7)},
		{VísperaInmaculadaConcepción, 2017, d(2017, 12, 7), d(2017, 12, 7)},
		{VísperaInmaculadaConcepción, 2018, d(2018, 12, 7), d(2018, 12, 7)},
		{VísperaInmaculadaConcepción, 2019, d(2019, 12, 7), d(2019, 12, 7)},
		{VísperaInmaculadaConcepción, 2020, d(2020, 12, 7), d(2020, 12, 7)},
		{VísperaInmaculadaConcepción, 2021, d(2021, 12, 7), d(2021, 12, 7)},
		{VísperaInmaculadaConcepción, 2022, d(2022, 12, 7), d(2022, 12, 7)},

		{InmaculadaConcepción, 2015, d(2015, 12, 8), d(2015, 12, 8)},
		{InmaculadaConcepción, 2016, d(2016, 12, 8), d(2016, 12, 8)},
		{InmaculadaConcepción, 2017, d(2017, 12, 8), d(2017, 12, 8)},
		{InmaculadaConcepción, 2018, d(2018, 12, 8), d(2018, 12, 8)},
		{InmaculadaConcepción, 2019, d(2019, 12, 8), d(2019, 12, 8)},
		{InmaculadaConcepción, 2020, d(2020, 12, 8), d(2020, 12, 8)},
		{InmaculadaConcepción, 2021, d(2021, 12, 8), d(2021, 12, 8)},
		{InmaculadaConcepción, 2022, d(2022, 12, 8), d(2022, 12, 8)},

		{Nochebuena, 2015, d(2015, 12, 24), d(2015, 12, 24)},
		{Nochebuena, 2016, d(2016, 12, 24), d(2016, 12, 24)},
		{Nochebuena, 2017, d(2017, 12, 24), d(2017, 12, 24)},
		{Nochebuena, 2018, d(2018, 12, 24), d(2018, 12, 24)},
		{Nochebuena, 2019, d(2019, 12, 24), d(2019, 12, 24)},
		{Nochebuena, 2020, d(2020, 12, 24), d(2020, 12, 24)},
		{Nochebuena, 2021, d(2021, 12, 24), d(2021, 12, 24)},
		{Nochebuena, 2022, d(2022, 12, 24), d(2022, 12, 24)},

		{Navidad, 2015, d(2015, 12, 25), d(2015, 12, 25)},
		{Navidad, 2016, d(2016, 12, 25), d(2016, 12, 25)},
		{Navidad, 2017, d(2017, 12, 25), d(2017, 12, 25)},
		{Navidad, 2018, d(2018, 12, 25), d(2018, 12, 25)},
		{Navidad, 2019, d(2019, 12, 25), d(2019, 12, 25)},
		{Navidad, 2020, d(2020, 12, 25), d(2020, 12, 25)},
		{Navidad, 2021, d(2021, 12, 25), d(2021, 12, 25)},
		{Navidad, 2022, d(2022, 12, 25), d(2022, 12, 25)},

		{AñoViejo, 2015, d(2015, 12, 31), d(2015, 12, 31)},
		{AñoViejo, 2016, d(2016, 12, 31), d(2016, 12, 31)},
		{AñoViejo, 2017, d(2017, 12, 31), d(2017, 12, 31)},
		{AñoViejo, 2018, d(2018, 12, 31), d(2018, 12, 31)},
		{AñoViejo, 2019, d(2019, 12, 31), d(2019, 12, 31)},
		{AñoViejo, 2020, d(2020, 12, 31), d(2020, 12, 31)},
		{AñoViejo, 2021, d(2021, 12, 31), d(2021, 12, 31)},
		{AñoViejo, 2022, d(2022, 12, 31), d(2022, 12, 31)},
	}

	for _, test := range tests {
		gotAct, gotObs := test.h.Calc(test.y)
		if !gotAct.Equal(test.wantAct) {
			t.Errorf("%s %d: got actual: %s, want: %s", test.h.Name, test.y, gotAct.String(), test.wantAct.String())
		}
		if !gotObs.Equal(test.wantObs) {
			t.Errorf("%s %d: got observed: %s, want: %s", test.h.Name, test.y, gotObs.String(), test.wantObs.String())
		}
	}
}
