class A
{
	int a=10;
	int b=20;
} 
class B extends A
{
	void add()  //Method to add 2 numbers
	{
	  int c=a+b;// Here we are using properties of class A
      System.out.println("The sum of a+b i.e 2 numbers are:"+c);
	}
}
class C extends A
{
	int c=30;
	void add()  //method to add three numbers
	{   
		int c=50;
		int d;
		d=(a+b+c);// Here also we are using properties a,b of class A
		System.out.println("The sum of 3 numbers is:"+d);
	}

}

class Hierarchical1
{
	public static void main(String[] args) 
	{
		System.out.println("Heirachical inheritance");//Exteding the properties and methods from one class to another is called inheritance
		//if same properties and methods or any are derived into more than one class it is called hierarchical inheritance
		B b =new B();
		b.add();
		C c=new C();
		c.add();
		
	}
}

