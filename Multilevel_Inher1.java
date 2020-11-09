class A
{
	int a=20;
	int b=40;
	void print()
	{
		System.out.println("Class A");
	}
} 
class B extends A
{   
	int c=a+b;
	void add()
	{
	  //int c=a+b;
      System.out.println("The sum of a+b are:"+c);
	}
	//System.out.println(" c value again is :"+c);
}
class C extends B//Deriving properties or methods or both from aleady exeded class B is called multi level inheritance
{
	int d;
	d=a+b+c;
	void display()
	{
	 System.out.println("Adding 70 to the a,b of  already exteded class B",+d);
	}
}
class Multilevel_Inher1
{
	public static void main(String[] args) 
	{
		System.out.println("Multiple inheritance");//Exteding the properties and methods from one class to another is called inheritance
		//if this is done from only one class to anohter it is called single inhertiance
		//B b=new B();
		//b.add();
		C c=new C();
		c.display();
	}
}

