package com.example.demo.service;

import com.example.demo.domain.Server;
import com.example.demo.exception.InvalidServerReferenceException;
import com.example.demo.repository.ServerRepository;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class ServerService {

    private static final String ERROR_MSG = "Server with id: %s does not exist";

    private final ServerRepository repository;

    public ServerService(ServerRepository repository) {
        this.repository = repository;
    }

    public List<Server> getAll() {
        return repository.findAll();
    }

    public Server findById(Long id) {
        return repository.findById(id)
                .orElseThrow(() -> new InvalidServerReferenceException(String.format(ERROR_MSG, id)));
    }

    public Server save(Server server) {
        return repository.save(server);
    }

    public void delete(Long id) {
        Server server = findById(id);
        server.setActive(false);
        repository.save(server);
    }
}
