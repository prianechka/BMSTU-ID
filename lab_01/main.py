from distutils.command.build_scripts import first_line_re
from install import check_hash


if __name__ == "__main__":
    try:
        print("Программа по выводу первых 10 чисел Фибоначчи!")
        print("Но сначала надо удостовериться, что у вас есть лицензия...\n\n")
        is_have_license = check_hash()

        if is_have_license:
            print("Вы - счастливый обладатель лицензии!")
            print("Числа Фибоначчи:")
            prev, current = 0, 1
            print(prev, current, end = " ")
            for i in range(8):
                prev, current = current, prev + current
                print(current, end = " ")
            print("\n\nДо встречи!")
        else:
            print("К сожалению, у вас нет лицензии!")
    except:
        print("Запустите exe с правами суперюзера!")