package factory

import (
	"errors"
	"fmt"
)

/*
简单工厂模型
用单独的类来常见这个实例
*/

type Drinks interface {
	Material() []string
	Price() int
}

//实例
type AmericanoCafe struct {
}

func (ac *AmericanoCafe) Material() []string {
	return []string{"water", "cafe"}
}
func (ac *AmericanoCafe) Price() int {
	return 16
}

type OrangeJuice struct {
}

func (oj *OrangeJuice) Material() []string {
	return []string{"water", "orange"}
}
func (oj *OrangeJuice) Price() int {
	return 14
}

/*
工厂类 GetDrinks 根据饮料名称返回饮料
*/
type DrinksFactory struct {
}

func (df *DrinksFactory) GetDrinks(drankName string) (Drinks, error) {
	switch drankName {
	case "americano_cafe":
		return &AmericanoCafe{}, nil
	case "orange_juice":
		return &OrangeJuice{}, nil
	}
	return nil, errors.New("illegal input")
}

/*
扩展方法：
1.实现Drinks方法
2.DrinksFactory.GetDrinks 添加判断。
*/

func StaticFactoryMethod() {
	factory := DrinksFactory{}
	drank, err := factory.GetDrinks("americano_cafe")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(drank.Material())
}
