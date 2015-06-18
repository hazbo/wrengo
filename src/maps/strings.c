// wrengo v0.0.1-dev
//
// (c) Harry Lawrence 2015
//
// @package wrengo
// @version 0.0.1-dev
//
// @author Harry Lawrence <http://github.com/hazbo>
//
// License: MIT
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

#include <string.h>
#include "strings.h"

int is_strings_contains(const char* className, const char* signature) {
  if ( strcmp( className, "Strings" ) == 0 ) {
    if ( strcmp( signature, "contains(_,_)" ) == 0 ) {
      return 1;
    }
  }
  return 0;
}
