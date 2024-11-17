// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract KIManagement {
    // Struct to store document details
    struct Document {
        string hash;
        bool isVerified;
    }

    // Struct to store Key Individual details
    struct KeyIndividual {
        string name;
        string qualification;
        string licenseType;
        string experience;
        mapping(uint256 => Document) documents;
        uint256 documentCount;
        bool isVerified;
    }

    // Mapping to store Key Individuals
    mapping(address => KeyIndividual) public keyIndividuals;
    address public admin;

    // Events
    event KIRegistered(address indexed kiAddress, string name);
    event DocumentUploaded(address indexed kiAddress, uint256 documentIndex, string documentHash);
    event DocumentVerified(address indexed kiAddress, uint256 documentIndex);
    event KIVerified(address indexed kiAddress);

    // Modifier to restrict access to admin only
    modifier onlyAdmin() {
        require(msg.sender == admin, "Only admin can perform this action");
        _;
    }

    // Constructor to set the admin
    constructor() {
        admin = msg.sender;
    }

    // Function to register a new Key Individual
    function registerKI(
        string memory _name,
        string memory _qualification,
        string memory _licenseType,
        string memory _experience
    ) public {
        require(bytes(keyIndividuals[msg.sender].name).length == 0, "KI already registered");
        
        KeyIndividual storage ki = keyIndividuals[msg.sender];
        ki.name = _name;
        ki.qualification = _qualification;
        ki.licenseType = _licenseType;
        ki.experience = _experience;
        ki.documentCount = 0;
        ki.isVerified = false;

        emit KIRegistered(msg.sender, _name);
    }

    // Function to upload a document
    function uploadDocument(string memory _documentHash) public {
        KeyIndividual storage ki = keyIndividuals[msg.sender];
        require(bytes(ki.name).length != 0, "KI not registered");

        ki.documents[ki.documentCount] = Document({
            hash: _documentHash,
            isVerified: false
        });
        emit DocumentUploaded(msg.sender, ki.documentCount, _documentHash);
        ki.documentCount++;
    }

    // Function for admin to verify a document
    function verifyDocument(address _kiAddress, uint256 _documentIndex) public onlyAdmin {
        KeyIndividual storage ki = keyIndividuals[_kiAddress];
        require(bytes(ki.name).length != 0, "KI not registered");
        require(_documentIndex < ki.documentCount, "Invalid document index");

        ki.documents[_documentIndex].isVerified = true;
        emit DocumentVerified(_kiAddress, _documentIndex);

        // Check if all documents are verified
        bool allVerified = true;
        for (uint256 i = 0; i < ki.documentCount; i++) {
            if (!ki.documents[i].isVerified) {
                allVerified = false;
                break;
            }
        }

        if (allVerified) {
            ki.isVerified = true;
            emit KIVerified(_kiAddress);
        }
    }

    // Function to get Key Individual details
    function getKIDetails(address _kiAddress) public view returns (
        string memory name,
        string memory qualification,
        string memory licenseType,
        string memory experience,
        uint256 documentCount,
        bool isVerified
    ) {
        KeyIndividual storage ki = keyIndividuals[_kiAddress];
        return (ki.name, ki.qualification, ki.licenseType, ki.experience, ki.documentCount, ki.isVerified);
    }

    // Function to get details of a specific document
    function getDocumentDetails(address _kiAddress, uint256 _documentIndex) public view returns (
        string memory hash,
        bool isVerified
    ) {
        KeyIndividual storage ki = keyIndividuals[_kiAddress];
        require(_documentIndex < ki.documentCount, "Invalid document index");
        Document storage doc = ki.documents[_documentIndex];
        return (doc.hash, doc.isVerified);
    }
}
