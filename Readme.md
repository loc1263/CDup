<style>
    .custom-font {
        font-size: 10px; /* Tamaño de la fuente en píxeles */
    }
</style>

# Dataset  

<pre>
<span classname="custom-font">

ID    Intereses         Sueldo             Saldo             Bonificación      Fecha  
0001  000001234567890   000000025000       00000000100000    00000005000000    20220429  
0002  000002345678901   000000035000       00000000150000    00000007500000    20220429  
0003  000003456789012   000000045000       00000000200000    00000010000000    20220429  
0004  000004567890123   000000055000       00000000250000    00000012500000    20220429  
0005  000005678901234   000000065000       00000000300000    00000015000000    20220429  
0006  000006789012345   000000075000       00000000350000    00000017500000    20220429  
0007  000007890123456   000000085000       00000000400000    00000020000000    20220429  
0008  <span style='color: red;'>000008901234567</span>   000000095000       00000000450000    00000022500000    20220429  
0009  000009012345678   000000105000       00000000500000    00000025000000    20220429  
0010  000010123456789   000000115000       00000000550000    00000027500000    20220429  
0011  000002345678901   000000035000       00000000150000    00000007500000    20220429  
0012  000006789012345   000000075000       00000000350000    00000017500000    20220429  
0013  000005678901234   000000065000       00000000300000    00000015000000    20220429  
0014  <span style='color: red;'>000008901234567</span>   000000095000       00000000450000    00000022500000    20220429  
0015  <span style='color: red;'>000008901234567</span>   000000096000       00000000450000    00000022500000    20220429  

</span>
</pre>

# Resultado  

``` go
go run .\CDup.go .\dataset.dat 6 21
Valor duplicado: 000002345678901 (líneas [3] y 12)
Valor duplicado: 000006789012345 (líneas [7] y 13)
Valor duplicado: 000005678901234 (líneas [6] y 14)
Valor duplicado: 000008901234567 (líneas [9] y 15)
Valor duplicado: 000008901234567 (líneas [9 15] y 16)
```
