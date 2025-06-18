package com.example.demo.exception;

public class InvalidServerReferenceException extends RuntimeException {

    public InvalidServerReferenceException(String message) {
        super(message);
    }
}
