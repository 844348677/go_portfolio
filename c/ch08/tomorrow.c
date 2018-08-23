#include <stdio.h>
#include <stdbool.h>

struct date
{
	int month;
	int day;
	int year;
};

int main (void)
{
	struct date today, tomorrow;
	int numberOfDays (struct date d);

	printf("Enter todays's date (mm dd yyyy): ");
	scanf("%i%i%i",&today.month,&today.day, &today.year);

	if(today.day != numberOfDays (today)){
		tomorrow.day = today.day + 1;
		tomorrow.month = today.month;
		tomorrow.year = today.year;
	}
	else if (today.month == 12){
		tomorrow.day = 1;
		tomorrow.month = 1;
		tomorrow.year = today.year + 1;
	}
	else{
		tomorrow.day = 1;
		tomorrow.month = today.month + 1;
		tomorrow.year = today.year;
	}

	printf ("tomorrow's date is %i/%i/%.2i.\n",tomorrow.month,tomorrow.day,tomorrow.year % 100);

}

int numberOfDays (struct date d)
{
	int days;
	bool isLeapYear (struct date d);

	const int daysPerMonth[12] = {31,28,31,30,31,30,31,31,30,31,30,31};

	if(isLeapYear (d) == true && d.month == 2)
		days = 29;
	else
		days = daysPerMonth[d.month - 1];

	return days;
}

bool isLeapYear (struct date d)
{
	bool leapYearFlag;

	if ( (d.year % 4 ==0 && d.year % 100 != 0) || d.year % 400 == 0)
		leapYearFlag = true;
	else 
		leapYearFlag = false;

	return leapYearFlag;
}