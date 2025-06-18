package com.example.demo.controller;

import com.example.demo.conf.EndpointConfiguration;
import com.example.demo.converter.ServerConverter;
import com.example.demo.domain.Server;
import com.example.demo.domain.OS;
import com.example.demo.dto.ServerDTO;
import com.example.demo.util.RequestCounter;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.ArrayList;

@RestController
@RequestMapping(EndpointConfiguration.SERVER_BASE_URL)
public class ServerController {

    private static final Logger LOGGER = LoggerFactory.getLogger(ServerController.class);

    private static final String LOG_GET_ALL = "action=getAllServers user=%s times=%s";
    private static final String LOG_GET_BY_ID = "action=getServerById id=%s user=%s times=%s";
    private static final String LOG_SAVE = "action=saveServer user=%s times=%s";
    private static final String LOG_REMOVE = "action=removeServer id=%s user=%s times=%s";

    private static final String SERVICE_NAME = "servers";
    private static final String DEFAULT_USER = "public";

    private final RequestCounter counter;

    private static final List<Server> servers = new ArrayList<Server>() {{
        add(new Server(0L, "database", "db.aws.com", OS.LINUX, true));
        add(new Server(1L, "sftp", "sftp.aws.com", OS.LINUX, true));
        add(new Server(2L, "datalake", "s3.aws.com", OS.LINUX, true));
    }};

    public ServerController(RequestCounter counter) {
        this.counter = counter;
    }

    /**
     * GET /api/server
     *
     * @return all available servers
     */
    @GetMapping
    public ResponseEntity<List<ServerDTO>> getAll() {
        String logContent = String.format(LOG_GET_ALL, DEFAULT_USER, counter.get(EndpointConfiguration.SERVER_BASE_URL));
        LOGGER.info(logContent);
        return new ResponseEntity<>(
                ServerConverter.fromEntityList(servers, ServerConverter::fromEntity),
                HttpStatus.OK
        );
    }
}
