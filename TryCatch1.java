class TryCatch1
{
	public static void main(String[] args) 
	{
		System.out.println("Hello welcome to java exception handling to uderstand them better ");
		System.out.println("Example for try catch finally in Arithematic Exception");
		int a = 20;
		int b = 10;
		int c = 10;// Program for try catch finally
		try
		{
		  int d = a/(b-c);//Dividing an inter by zero abnormal condition
		  System.out.println("d="+d);
		}
		catch(ArithmeticException e)  //catching the exception thrown by try block
		{
			e.printStackTrace();
			System.out.println(e);
			System.out.println(e.getMessage());
		}

		//int e = a/(b+c);
		//System.out.println(" The e value is :"+e);
		finally
		{
			System.out.println("This is finally block");
		}
	}
}

