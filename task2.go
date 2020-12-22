"""
Задание 2.

Реализуйте два алгоритма.

Первый, в виде функции, должен обеспечивать поиск минимального значения для списка.
В основе алгоритма должно быть сравнение каждого числа со всеми другими элементами списка.
Сложность такого алгоритма: O(n^2) - квадратичная.


Второй, в виде функции, должен обеспечивать поиск минимального значения для списка.
Сложность такого алгоритма: O(n) - линейная.

Примечание:
Построить список можно через генератор списка.
Если у вас возникают сложности, постарайтесь подумать как можно решить задачу,
а не писать "мы это не проходили)".
Алгоритмизатор должен развивать мышление, а это прежде всего практика.
А без столкновения со сложностями его не развить.
"""
arr = []
for i in range(5, 15):
    arr.append(i)
print(arr)


## O(N)

def min_v(arr):
    min = arr[0]
    for i in range(len(arr)):
        if min > arr[i]:
            min = arr[i]

    return min


# O(N^2)
def min_v2(arr):
    min = arr[0]
    for z in range(len(arr)):
        for j in range(len(arr)):
            if min > arr[j]:
                min = arr[j]

    return min


print(min_v(arr))
print(min_v2(arr))
