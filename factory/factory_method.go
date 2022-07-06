package factory

import "fmt"

/*
工厂方法模型
定义一个用于创建对象接口，让子类决定实例化哪一个类
工厂方法使一个类的实例化延迟到其子类
结构体依赖于 static _factory_method
*/

//工厂抽象类(接口)
type Factory interface {
	CreateDrinkFactory() Drinks
}

//咖啡工厂
type CoffeeFactory struct {
}

func (cf *CoffeeFactory) CreateDrinkFactory() Drinks {
	return &AmericanoCafe{}
}

//果汁工厂
type FruitJuiceFactory struct {
}

func (cf *FruitJuiceFactory) CreateDrinkFactory() Drinks {
	return &OrangeJuice{}
}

/*
这样咖啡工厂和果汁工厂就会被分成两大类
对比简单工厂和工厂方法：
简单工厂模型的最大优点在于工厂类中包含了必要的判断逻辑，根据客户端的选择条件动摇示例话相关的类
对于客户端来说，去除了与具体产品的依赖，但是违背了开放-封闭原则

工厂方法模型实现时，客户端需要决定实例化哪一个工厂类，选择判断的问题还是存在，只是把简单工厂内部的判断逻辑移到客户端代码来执行
但是很好的解决了违背开放-封闭原则的现象。
*/

//将判断下沉至子类
func FactoryMethod() {
	/*我要一杯咖啡*/
	factory := CoffeeFactory{}
	drank := factory.CreateDrinkFactory()
	fmt.Println(drank.Material())
}
