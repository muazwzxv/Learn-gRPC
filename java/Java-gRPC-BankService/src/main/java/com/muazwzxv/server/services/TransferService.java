package com.muazwzxv.server.services;

import com.muazwzxv.models.TransferRequest;
import com.muazwzxv.models.TransferResponse;
import com.muazwzxv.models.TransferServiceGrpc;
import com.muazwzxv.server.observer.TransferStreamObserver;
import io.grpc.stub.StreamObserver;

public class TransferService extends TransferServiceGrpc.TransferServiceImplBase {

    @Override
    public StreamObserver<TransferRequest> transfer(StreamObserver<TransferResponse> responseObserver) {
        return new TransferStreamObserver(responseObserver);
    }
}
