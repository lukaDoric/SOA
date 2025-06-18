package com.example.demo.util;

import org.springframework.stereotype.Component;

import java.util.HashMap;
import java.util.Map;

@Component
public class RequestCounter {

    private Map<String, Long> counter;

    public RequestCounter() {
        counter = new HashMap<>();
    }

    public Long get(String endpoint) {
        counter.put(endpoint, counter.getOrDefault(endpoint, 0L) + 1);
        return counter.get(endpoint);
    }
}
