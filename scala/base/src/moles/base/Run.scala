package moles.base

object Run {
  def main(args: Array[String]) = {
    val i = 10;
    val s = "Hello, Scala!";
    println("Hello, world");
    var a=11;
    var b=5;
    printf("%d+%d=%d\n",a,b,add(a,b));
    val info=new Info("test",1);
    println(info);
    info.print();
    println("factorial "+factorial(5));
  }

  def add( a:Int, b:Int ) : Int = {
    return a+b
  }

  def factorial(n: Int): Int= {
    if(n<0){
      throw new IllegalArgumentException();
    }
    if(n<2){
      return 1;
    }
    return factorial(n-1)*n;
  }
}
