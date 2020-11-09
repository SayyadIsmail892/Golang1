class ArrayIndex1
{
	public static void main(String[] args) 
	{
		System.out.println("Arrat index out of bounds exception ");
		int [] a={1,2,3,4,5};
		System.out.println(a[10]);//Trying to access an integer out of bounds causes an array out of bound exception
		System.out.Println("After exception handling");//This line will not printed once exeption occurs it terminates the program abnormally
		//for this line to execute  we have to use exception hadling mechanisms like try catch
	}
}
