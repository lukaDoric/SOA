package com.example.demo.dto;

import com.example.demo.domain.OS;
import lombok.*;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class ServerDTO {

    private Long id;
    private String hostname;
    private String domain;
    private OS os;
}
