#include <stdio.h>

struct entry
{
	char word[15];
	char definition[50];
};

int compareString(const char s1[], const char s2[])
{
	int i = 0, answer;
	while (s1[i] == s2[i] && s1[i] != '\0' && s2[i] != '\0')
		++i;

	if (s1[i] < s2[i])
		answer = -1;	// s1 < s2
	else if (s1[i] == s2[i])
		answer = 0;
	else
		answer = 1;

	return answer;
}

