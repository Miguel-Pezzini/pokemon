#ifndef COMBAT_H
#define COMBAT_H
#include <vector>
#include "menu.h" 
#include "routes.h"
#include <cstdlib>  // Para rand() e srand()
#include <ctime>    // Para time()

using namespace std;
bool inCombatInBush(Character &character, int route);

#endif