package method

// Blir problem med den krävda minnesmöngden . Behöver spara ca 2**59 intar
// vilket borde bli 2 ** 59  * 8 bytes
// 2 ** 10 ca 1000 fast lite över
//  vi får över  10**(3*6) bytes att allocera
// k 3 , M , 6 , G 9 , T 12 , men vi ska till 18
// slutsats statespace är för stort för att kunna utföra value iteration i den klassicka meningen
