class A
{
	int a=10;
	int b=20;
} 
class B extends A
{
	void add()
	{
	  int c=a+b;
      System.out.println("The sum of a+b are:"+c);
	}
}
class Single_Inher1 
{
	public static void main(String[] args) 
	{
		System.out.println("Single inheritance");//Exteding the properties and methods from one class to another is called inheritance
		//if this is done from only one class to anohter it is called single inhertiance
		B b=new B();
		b.add();
		
	}
}
