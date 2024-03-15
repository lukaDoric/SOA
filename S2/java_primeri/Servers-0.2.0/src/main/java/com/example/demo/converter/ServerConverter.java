package com.example.demo.converter;

import com.example.demo.domain.Server;
import com.example.demo.dto.ServerDTO;

public class ServerConverter extends AbstractConverter {

    public static ServerDTO fromEntity(Server server) {
        return ServerDTO.builder()
                .id(server.getId())
                .hostname(server.getHostname())
                .domain(server.getDomain())
                .os(server.getOs())
                .build();
    }

    public static Server toEntity(ServerDTO serverDTO) {
        return Server.builder()
                .hostname(serverDTO.getHostname())
                .domain(serverDTO.getDomain())
                .os(serverDTO.getOs())
                .active(true)
                .build();
    }
}
