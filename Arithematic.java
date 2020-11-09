class Arithematic 
{
	public static void main(String[] args) 
	{
		System.out.println("Hello welcome to java exception class to uderstand them better ");
		System.out.println("Example for Arithematic Exception (part of checked exception)");
		int a = 20;
		int b = 10;
		int c = 10;
		int d = a/(b-c);//Dividing an inter by zero
		System.out.println("d="+d);
		int e = a/(b+c);
		System.out.println(" The e value is :"+e);
	}
}
