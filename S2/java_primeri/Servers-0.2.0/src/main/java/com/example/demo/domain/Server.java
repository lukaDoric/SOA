package com.example.demo.domain;

import lombok.*;
import org.hibernate.annotations.Where;

import javax.persistence.*;

@Entity
@Where(clause = "active=true")
@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class Server {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;

    @Column(nullable = false)
    private String hostname;

    @Column(nullable = false)
    private String domain;

    @Column(nullable = false)
    @Enumerated(EnumType.ORDINAL)
    private OS os;

    @Column(nullable = false)
    private boolean active;
}
