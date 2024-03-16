package com.example.demo.domain;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

@NoArgsConstructor
@AllArgsConstructor
@Getter
@Setter
public class Log {

    private static final String DELIMITER = "|";
    private static final String FORMAT = "%s" + DELIMITER + "%s" + DELIMITER + "%s";
    private static final DateTimeFormatter DATE_FORMATTER = DateTimeFormatter.ISO_LOCAL_DATE_TIME;

    private LocalDateTime timestamp;
    private String service;
    private String content;

    public Log(String service, String content) {
        this.timestamp = LocalDateTime.now();
        this.service = service;
        this.content = content;
    }

    @Override
    public String toString() {
        return String.format(FORMAT, DATE_FORMATTER.format(timestamp), service, content);
    }
}
