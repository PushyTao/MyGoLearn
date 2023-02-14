1. Go语言中的数组初始化方法比较多，支持根据参数的数量来确定数组的大小，支持对某一个位置的参数进行设定

   <img src="./Collections.assets/image-20230213214202333.png" alt="image-20230213214202333" align="left" style="zoom:50%;" />

2. 数组的遍历方式类似于C++中遍历std::vector<T> vet;

   可以用for(auto at: vet) 也可以for i通过下标来进行访问

3. 切片的访问可以直接通过有效下标来访问任意的一个元素，也可以通过`[start, end]`来访问某一段的内容

   ![image-20230214223017582](./Collections.assets/image-20230214223017582.png)

   这样的方式可以直接用来删减一个或者是几个数据

4. 