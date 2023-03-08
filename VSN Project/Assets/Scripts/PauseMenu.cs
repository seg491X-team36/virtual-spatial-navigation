using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using UnityEngine.SceneManagement;

public class PauseMenu : MonoBehaviour
{

    public static bool isPaused = false;
    public GameObject pauseMenuUI;

    // Update is called once per frame
    void Update()
    {
        if(Input.GetKeyDown(KeyCode.Escape)){
            if (isPaused){
                Resume();
            } else {
                Pause();
            }
        }
    }

    public void Resume(){

        Cursor.lockState = CursorLockMode.Locked;
        pauseMenuUI.SetActive(false);
        Time.timeScale = 1f;
        isPaused = false;

    }

    public void Pause(){

        Cursor.lockState = CursorLockMode.Confined;
        pauseMenuUI.SetActive(true);
        Time.timeScale = 0f;
        isPaused = true;

    }

    // public void LoadControls(){
    //     SceneManager.LoadScene("Main Menu");
    //     //Replace with variable for easier access
    // }

    public void QuitExperiment(){
        SceneManager.LoadScene("Main Menu");
        //Replace with variable for easier access
    }
}
