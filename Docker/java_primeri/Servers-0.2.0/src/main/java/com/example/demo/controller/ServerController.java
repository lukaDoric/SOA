package com.example.demo.controller;

import com.example.demo.conf.EndpointConfiguration;
import com.example.demo.converter.ServerConverter;
import com.example.demo.domain.Log;
import com.example.demo.dto.ServerDTO;
import com.example.demo.service.ServerService;
import com.example.demo.util.RequestCounter;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;
import java.util.List;
import java.util.concurrent.TimeoutException;

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

    private final ServerService service;
    private final RequestCounter counter;

    public ServerController(ServerService service, RequestCounter counter) {
        this.service = service;
        this.counter = counter;
    }

    /**
     * GET /api/server
     *
     * @return all available servers
     */
    @GetMapping
    public ResponseEntity<List<ServerDTO>> getAll() throws IOException, TimeoutException {
        String logContent = String.format(LOG_GET_ALL, DEFAULT_USER, counter.get(EndpointConfiguration.SERVER_BASE_URL));
        LOGGER.info(logContent);
        return new ResponseEntity<>(
                ServerConverter.fromEntityList(service.getAll(), ServerConverter::fromEntity),
                HttpStatus.OK
        );
    }

    /**
     * GET /api/server/{id}
     *
     * @param id of requested server
     * @return server with requested id
     */
    @GetMapping(EndpointConfiguration.SERVER_ID_ENDPOINT)
    public ResponseEntity<ServerDTO> getById(@PathVariable Long id) throws IOException, TimeoutException {
        String logContent = String.format(LOG_GET_BY_ID, id, DEFAULT_USER, counter.get(EndpointConfiguration.SERVER_ID_URL));
        LOGGER.info(logContent);
        return new ResponseEntity<>(
                ServerConverter.fromEntity(service.findById(id)),
                HttpStatus.FOUND
        );
    }

    /**
     * POST /api/server
     *
     * @param server that needs to be saved
     * @return added server
     */
    @PostMapping
    public ResponseEntity<ServerDTO> save(@RequestBody ServerDTO server) throws IOException, TimeoutException {
        String logContent = String.format(LOG_SAVE, DEFAULT_USER, counter.get(EndpointConfiguration.SERVER_BASE_URL));
        LOGGER.info(logContent);
        return new ResponseEntity<>(
                ServerConverter.fromEntity(service.save(ServerConverter.toEntity(server))),
                HttpStatus.OK
        );
    }

    /**
     * DELETE /api/server/{id}
     *
     * @param id of vehicle that needs to be deleted
     * @return message about action results
     */
    @DeleteMapping(value = EndpointConfiguration.SERVER_ID_ENDPOINT, produces = MediaType.TEXT_PLAIN_VALUE)
    public ResponseEntity<String> remove(@PathVariable Long id) throws IOException, TimeoutException {
        String logContent = String.format(LOG_REMOVE, id, DEFAULT_USER, counter.get(EndpointConfiguration.SERVER_BASE_URL));
        LOGGER.info(logContent);
        service.delete(id);
        return new ResponseEntity<>("Server successfully deleted!", HttpStatus.OK);
    }
}
