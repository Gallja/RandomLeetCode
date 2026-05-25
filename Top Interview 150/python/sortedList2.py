from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def scorri_lista(testa: Optional[ListNode]):
    corrente = testa
    
    while corrente is not None:
        print(corrente.val, end=" -> ")
        
        corrente = corrente.next
        
    print("None")

def merge_and_sort(testa_1: Optional[ListNode], testa_2: Optional[ListNode]) -> Optional[ListNode]:
    corrente_1 = testa_1
    corrente_2 = testa_2
    
    cnt_1 = 0
    cnt_2 = 0

    while corrente_1 is not None:
        cnt_1 += 1
        corrente_1 = corrente_1.next

    while corrente_2 is not None:
        cnt_2 += 1
        corrente_2 = corrente_2.next 

    cnt_fin = cnt_1 + cnt_2

    corrente_1 = testa_1
    corrente_2 = testa_2

    if corrente_1.val <= corrente_2.val:
        testa_3 = ListNode(corrente_1.val)
        corrente_1 = corrente_1.next
    else:
        testa_3 = ListNode(corrente_2.val)
        corrente_2 = corrente_2.next

    corrente_3 = testa_3

    for i in range(cnt_fin):
        if corrente_1 is None:
            while corrente_2 is not None:
                corrente_3.next = corrente_2
                corrente_2 = corrente_2.next 
            break

        if corrente_2 is None:
            while corrente_1 is not None:
                corrente_3.next = corrente_1
                corrente_1 = corrente_1.next
            break

        if corrente_1.val <= corrente_2.val:
            corrente_3.next = corrente_1
            corrente_1 = corrente_1.next
        else:
            corrente_3.next = corrente_2
            corrente_2 = corrente_2.next

        corrente_3 = corrente_3.next

    return testa_3
                   

nodo_1_1 = ListNode(1)
nodo_1_2 = ListNode(2)
nodo_1_3 = ListNode(4)

nodo_2_1 = ListNode(1)
nodo_2_2 = ListNode(3)
nodo_2_3 = ListNode(4)

nodo_1_1.next = nodo_1_2
nodo_1_2.next = nodo_1_3
testa_lista_1 = nodo_1_1

nodo_2_1.next = nodo_2_2
nodo_2_2.next = nodo_2_3
testa_lista_2 = nodo_2_1

print("Scorrimento della lista:")
scorri_lista(testa_lista_1)

print("\nScorrimento della lista:")
scorri_lista(testa_lista_2)

print("\nDopo l'ordinamento:")
testa_lista_3 = merge_and_sort(nodo_1_1, nodo_2_1)
scorri_lista(testa_lista_3)