package pswarm_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/rwcarlsen/optim/pswarm"
	"github.com/rwcarlsen/optim/pswarm/population"
)

func TestAckley(t *testing.T) {
	it := pswarm.SimpleIter{}
	ev := pswarm.SerialEvaler{}
	mv := &pswarm.SimpleMover{
		Cognition: pswarm.DefaultCognition,
		Social:    pswarm.DefaultSocial,
	}
	obj := pswarm.NewObjectivePrinter(pswarm.SimpleObjectiver(Ackley))

	lb := []float64{-5, -5}
	ub := []float64{5, 5}
	minv := []float64{0.5, 0.5}
	maxv := []float64{1.5, 1.5}

	rand.Seed(time.Now().Unix())
	pop := population.NewRandom(30, lb, ub, minv, maxv)

	var err error
	for i := 0; i < 500; i++ {
		pop, err = it.Iterate(pop, obj, ev, mv)
		if err != nil {
			t.Fatal(err)
		}
	}

	val, pos := pop.Best()
	t.Log("BestVal: ", val)
	t.Log("Found at: ", pos)
}

func Ackley(v []float64) float64 {
	x := v[0]
	y := v[1]
	return -20*math.Exp(-0.2*math.Sqrt(0.5*(x*x+y*y))) -
		math.Exp(0.5*(math.Cos(2*math.Pi*x)+math.Cos(2*math.Pi*y))) +
		20 + math.E
}