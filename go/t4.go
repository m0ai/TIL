package t4


type (
	Animal struct {
		species string
		age int
	}

	AnimalHouse struct {
		name string
		sizeInMeters int
	}

	AnimalFactory struct {
		species string
		houseName string
	}
)

func (af AnimalFactory) NewAnimal(age int) {
	return Animal{
		species: af.species,
		age:     age,
	}
}

func (af AnimalFactory) NewHouse(sizeInMeters int) {
	return Animal{
		name:         af.houseName,
		sizeInMeters: sizeInMeters,
	}
}

func main () {
	dogFactory := AnimalFactory{
		species: "dog",
		houseName: "kennel"
	}
	dog := dogFactory.NewAnimal(2)
	kennel := dogFactory.NewHouse(3)


	horseFactory := AnimalFactory{
		species: "horse",
		houseName: "stable"
	}

	horse := horseFactory.NewAnimal(12)
	stable := horseFactory.NewHouse(30)
}


