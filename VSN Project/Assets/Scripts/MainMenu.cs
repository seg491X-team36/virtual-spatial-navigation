using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;

public class MainMenu : MonoBehaviour
{
    public void StartExperiment(){
        SceneManager.LoadScene(SceneManager.GetActiveScene().buildIndex + 1);
        //Replace above logic with grabbing experiment information from Backend
    }

    public void QuitGame(){
        Debug.Log("Quit");
        Application.Quit();
    }
}
