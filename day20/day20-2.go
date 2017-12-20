package main

import (
	"os"
	"bufio"
	"regexp"
	"fmt"
	"strings"
	"strconv"
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
	enabled bool
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

		p := Particle{i, pdn[0], adn[0], vdn[0], pdn[1], adn[1], vdn[1], pdn[2], adn[2], vdn[2], true}
		parts = append(parts,p)
		i++
	}

	return parts
}

func (s Particles) Equal(i, j int) bool {
	return (s[i].x == s[j].x && s[i].y == s[j].y && s[i].z == s[j].z)
}

func distance(p Particle) int {
	return int(math.Abs(float64(p.x)) + math.Abs(float64(p.y)) + math.Abs(float64(p.z)))
}

func mark_collisions(particles Particles) int {
	new_collisions := 0

	for i := 0; i < len(particles); i++ {
		if !particles[i].enabled {
			continue
		}

		for j := 0; j < len(particles); j++ {
			if i == j || !particles[j].enabled {
				continue
			}
			if particles.Equal(i,j) {
				particles[i].enabled = false
				particles[j].enabled = false
				new_collisions++
			}
		}
	}
	return new_collisions
}

func get_num_enabled(particles Particles) int {
	enabled := 0
	for _, v := range(particles) {
		if v.enabled {
			enabled++
		}
	}
	return enabled
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

		collisions := mark_collisions(particles)
		fmt.Printf("On iteration %d, we removed an additional %d, leaving %d\n",i,collisions,get_num_enabled(particles))
	}
}

func main() {
	particles := read_file(os.Args[1])
	iterate(particles)
}
