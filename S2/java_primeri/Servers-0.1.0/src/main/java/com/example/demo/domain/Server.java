package com.example.demo.domain;

import lombok.*;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class Server {

    private Long id;
    private String hostname;
    private String domain;
    private OS os;
    private boolean active;
}
