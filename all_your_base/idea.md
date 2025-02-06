🛠️ Étapes de la méthode de Horner

Soit un nombre écrit en base XX sous la forme :
(anan−1...a2a1a0)X
(an​an−1​...a2​a1​a0​)X​

où chaque aiai​ est un chiffre en base XX.

Nous utilisons la factorisation suivante :
(((an×X+an−1)×X+an−2)×X+...+a1)×X+a0
(((an​×X+an−1​)×X+an−2​)×X+...+a1​)×X+a0​

Ceci permet de transformer la somme des puissances en une série de multiplications et d’additions successives, ce qui est plus rapide.
🎯 Exemple de conversion

Convertissons (2345)7(2345)7​ en base 10.

📌 Rappel de la conversion classique :
23457=2×73+3×72+4×71+5×70
23457​=2×73+3×72+4×71+5×70
=2×343+3×49+4×7+5
=2×343+3×49+4×7+5
=686+147+28+5=866
=686+147+28+5=866

📌 Avec Horner, on réécrit cela de façon factorisée :
(((2×7+3)×7+4)×7+5)
(((2×7+3)×7+4)×7+5)

Faisons les calculs étape par étape :

    2×7+3=14+3=172×7+3=14+3=17
    17×7+4=119+4=12317×7+4=119+4=123
    123×7+5=861+5=866123×7+5=861+5=866

✅ On retrouve bien (2345)7=(866)10(2345)7​=(866)10​.


Etape de la conversion vers une nouvelle base : 

1️⃣ Conversion de la base 10 vers une base YY (Partie entière)
Méthode des divisions successives

Elle consiste à :

    Diviser le nombre par la base YY.
    Récupérer le reste (c'est un chiffre du nombre dans la nouvelle base).
    Remplacer le nombre par le quotient et répéter jusqu'à obtenir un quotient de 0.
    Lire les restes à l’envers pour obtenir le nombre final.

🔍 Exemple

Convertissons (866)10​ en base 7.
Division	Quotient	Reste
866 ÷ 7	123	5
123 ÷ 7	17	4
17 ÷ 7	2	3
2 ÷ 7	0	2

Lecture des restes de bas en haut : (2345)7(2345)7​
✅ On retrouve bien notre nombre d'origine !
