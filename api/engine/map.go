package engine

// BoardMap represents the game Map
type BoardMap struct {
	areas map[string]*MapArea
	links map[string][]*MapArea
}

// AddNode adds an area to the graph
func (m *BoardMap) AddNode(a *MapArea) {
	if m.areas == nil {
		m.areas = make(map[string]*MapArea)
	}
	m.areas[a.name] = a
}

// AddLink adds a link between two areas
func (m *BoardMap) AddLink(a, b *MapArea, permission string) {
	if m.links == nil {
		m.links = make(map[string][]*MapArea)
	}
	m.links[a.name] = append(m.links[a.name], b)
	m.links[b.name] = append(m.links[b.name], a)

	//TODO: implement permissions
}

// MapArea represents an area in the map
type MapArea struct {
	name         string
	areaType     string
	units        []Unit
	flag         string
	factory      bool
	factoryBuild bool
}
