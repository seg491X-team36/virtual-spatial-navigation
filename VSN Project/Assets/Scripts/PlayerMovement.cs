using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class PlayerMovement : MonoBehaviour
{
    
    [SerializeField] float moveSpeed = 10f;
    [SerializeField] float rotateSpeedH = 5f;
    [SerializeField] float rotateSpeedV = 5f;
    [SerializeField] float jumpForce = 8f;
    [SerializeField] float gravity = 30f;

    private CharacterController player;
    private Vector3 moveDirection;
    private float yaw = 0f;
    private float pitch = 0f;

    void Start()
    {
        player = GetComponent<CharacterController>();
        Cursor.lockState = CursorLockMode.Locked;
        Cursor.visible = false;
    }

    void Update()
    {
        HandleMovementInput();
        HandleMouseLook();
        HandleJump();

        HandleFinalMovement();
    }

    private void HandleMovementInput()
    {
        float horizontalInput = Input.GetAxis("Horizontal") * moveSpeed;
        float verticalInput = Input.GetAxis("Vertical") * moveSpeed;

        float moveDirectionY = moveDirection.y;
        moveDirection = (transform.TransformDirection(Vector3.forward) * verticalInput) + (transform.TransformDirection(Vector3.right) * horizontalInput);
        moveDirection.y = moveDirectionY;
    }

    private void HandleMouseLook()
    {
        yaw += rotateSpeedH * Input.GetAxis("Mouse X");
        pitch -= rotateSpeedV * Input.GetAxis("Mouse Y");
        pitch = Mathf.Clamp(pitch, -80, 60);
        player.transform.eulerAngles = new Vector3(pitch, yaw, 0f);
    }

    private void HandleJump()
    {
        if (Input.GetButtonDown("Jump") && player.isGrounded)
            moveDirection.y = jumpForce;
    }

    private void HandleFinalMovement()
    {
        if (!player.isGrounded)
            moveDirection.y -= gravity * Time.deltaTime;
        player.Move(moveDirection * Time.deltaTime);
    }
}
