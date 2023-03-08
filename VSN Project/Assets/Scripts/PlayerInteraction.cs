using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class PlayerInteraction : MonoBehaviour
{
    int doorsOpened = 0;

    void Update()
    {
        if (Input.GetButtonDown("Interact"))
        {
            RaycastHit hit;
            Ray ray = Camera.main.ScreenPointToRay(Input.mousePosition);

            if (Physics.Raycast(ray, out hit, 5))
            {
                print(hit.collider);
                if (hit.rigidbody != null && hit.rigidbody.gameObject.layer == 6)
                {
                    if (doorsOpened++ < 1)
                    {
                        Destroy(hit.rigidbody.gameObject);
                    }  
                }
            }
        }
    }
}
