package factory

import "fmt"

/*
抽象工厂模式
部分结构体依赖与factory_method
*/

//其实是对工厂模式一种扩展 引入甜点
type Snack interface {
	GetName() string
	GetWeight() int
}

//实例
type CoCoCookie struct {
}

func (cc *CoCoCookie) GetName() string {
	return "coco"
}
func (cc *CoCoCookie) GetWeight() int {
	return 1
}

type FruitCookie struct {
}

func (cc *FruitCookie) GetName() string {
	return "fruit"
}
func (cc *FruitCookie) GetWeight() int {
	return 2
}

// 抽象工厂
type AbstractFactory interface {
	CreateDrink() Drinks
	CreateSnack() Snack
}

//咖啡工厂
type AbstractCoffeeFactory struct {
}

func (cf *AbstractCoffeeFactory) CreateDrink() Drinks {
	return &AmericanoCafe{}
}
func (cf *AbstractCoffeeFactory) CreateSnack() Snack {
	return &CoCoCookie{}
}

//果汁工厂
type AbstractFruitJuiceFactory struct {
}

func (cf *AbstractFruitJuiceFactory) CreateDrink() Drinks {
	return &OrangeJuice{}
}
func (cf *AbstractFruitJuiceFactory) CreateSnack() Snack {
	return &FruitCookie{}
}

func UseAbstractFactory() {
	drink := func(drinks Drinks) {
		fmt.Println(drinks.Material())
	}
	eat := func(snack Snack) {
		fmt.Println(snack.GetName())
	}

	/*
		这里可以用简单工厂进行switch选择，可以利用依赖注入等技术来选择生成什么点
		由于GO对接口的友好，这里用抽象工厂天然的很方便
	*/
	//水果店
	//shop := AbstractFruitJuiceFactory{}
	//咖啡店
	shop := AbstractCoffeeFactory{}
	drink(shop.CreateDrink())
	eat(shop.CreateSnack())
}
