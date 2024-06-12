#include <stdio.h>
#include <stdlib.h>

struct Data {
};

int main() {
    struct Data *ptr1;
    struct Data *ptr2;

    ptr1 = (struct Data *)malloc(sizeof(struct Data));
    ptr2 = (struct Data *)malloc(sizeof(struct Data));

    if (ptr1 == NULL || ptr2 == NULL) {
        printf("malloc error.\n");
        return 1;
    }

	if (ptr1 == ptr2) {
		printf("same address - ptr1=%p ptr2=%p\n", ptr1, ptr2);
	}else{
        printf("different address - ptr1=%p ptr2=%p\n", ptr1,ptr2);
    }

    free(ptr1);
    free(ptr2);

    return 0;
}
