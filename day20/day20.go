package main

import (
	"os"
	"bufio"
	"regexp"
	"fmt"
	"strings"
	"strconv"
	"sort"
	"math"
)

type Particle struct {
	id int
	x int
	xa int
	xv int
	y int
	ya int
	yv int
	z int
	za int
	zv int
}

type Particles []Particle

func read_file(f string) Particles {
	file, _ := os.Open(f)
	s := bufio.NewScanner(file)

	rmatch, _ := regexp.Compile(`p=<(.+)>, v=<(.+)>, a=<(.+)>`)

	parts := make([]Particle,0)

	i := 0
	for s.Scan() {
		line := s.Text()
		rm := rmatch.FindStringSubmatch(line)
		if rm != nil {
			fmt.Println("Invalidn line: ", line)
		}

		pd := strings.Split(rm[1],",")
		pdn := make([]int, 0)
		for _, v := range(pd) {
			pn, _ := strconv.Atoi(v)
			pdn = append(pdn, pn)
		}
		vd := strings.Split(rm[2], ",")
		vdn := make([]int, 0)
		for _, v := range(vd) {
			vn, _ := strconv.Atoi(v)
			vdn = append(vdn, vn)
		}

		ad := strings.Split(rm[3], ",")
		adn := make([]int, 0)
		for _, v := range(ad) {
			an, _ := strconv.Atoi(v)
			adn = append(adn, an)
		}

		p := Particle{i, pdn[0], adn[0], vdn[0], pdn[1], adn[1], vdn[1], pdn[2], adn[2], vdn[2]}
		parts = append(parts,p)
		i++
	}

	return parts
}

func (s Particles) Len() int {
	return len(s)
}
func (s Particles) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Particles) Less(i, j int) bool {
	return distance(s[i]) < distance(s[j])
}

func distance(p Particle) int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)) + math.Abs(float64(p.z)))
}

func iterate(particles Particles) {
	for i:= 0; i < 500; i++ {
		for j, v := range(particles) {
			v.xv += v.xa
			v.yv += v.ya
			v.zv += v.za
			v.x += v.xv
			v.y += v.yv
			v.z += v.zv
			particles[j] = v
		}

		// sort our particles based on their current distance
		sort.Sort(Particles(particles))

		// print the top 5 smallest
		fmt.Printf("On iteration %d the top 5 are: ", i)
		fmt.Printf("%d %d %d %d %d\n",particles[0].id, particles[1].id,particles[2].id,particles[3].id,particles[4].id)
	}
}

func main() {
	particles := read_file(os.Args[1])
	iterate(particles)
}
