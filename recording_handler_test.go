package main

import (
	assert "github.com/pilu/miniassert"
	"testing"
)

func TestRecordingHandler_WithExistingRecordingId(t *testing.T) {
	recorder := newTestRequest("GET", "/recordings/833f00e1-781f-4edd-90e4-e52712618862.json")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"id":"833f00e1-781f-4edd-90e4-e52712618862","name":"Get Lucky","comment":"","length":367000,"artists":[{"id":"056e4f3e-d505-4dad-8ec1-d04f521cbb56","name":"Daft Punk"},{"id":"149f91ef-1287-46da-9a8e-87fee02f1471","name":"Pharrell Williams"}]}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 200, recorder.Code)
}

func TestRecordingHandler_WithRecordingIdNotFound(t *testing.T) {
	recorder := newTestRequest("GET", "/recordings/00000000-0000-0000-0000-000000000000.json")

	body := string(recorder.Body.Bytes())
	assert.Equal(t, `{"error":"recording not found"}`+"\n", body)
	assert.Equal(t, 404, recorder.Code)
}

func TestRecordingHandler_WithInvalidRecordingId(t *testing.T) {
	recorder := newTestRequest("GET", "/recordings/bad-uuid.json")

	body := string(recorder.Body.Bytes())
	assert.Equal(t, "", body)
	assert.Equal(t, 400, recorder.Code)
}

func TestRecordingHandler_WithNonExistingReleaseIdAndExistingRecordingId(t *testing.T) {
	recorder := newTestRequest("GET", "/releases/00000000-0000-0000-0000-000000000000/recordings/833f00e1-781f-4edd-90e4-e52712618862.json")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"error":"release not found"}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 404, recorder.Code)
}

func TestRecordingHandler_WithWrongReleaseIdAndExistingRecordingId(t *testing.T) {
	// Release is "Harder, Better, Faster, Stronger" but recording is "Get Lucky" which is in "Random Access Memories"
	recorder := newTestRequest("GET", "/releases/e1ed2270-c44f-4c72-8836-140579b211fa/recordings/833f00e1-781f-4edd-90e4-e52712618862.json")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"error":"recording not found"}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 404, recorder.Code)
}

func TestRecordingHandler_WithExistingReleaseIdAndRecordingId(t *testing.T) {
	recorder := newTestRequest("GET", "/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings/833f00e1-781f-4edd-90e4-e52712618862.json")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"id":"833f00e1-781f-4edd-90e4-e52712618862","name":"Get Lucky","comment":"","length":367000,"artists":[{"id":"056e4f3e-d505-4dad-8ec1-d04f521cbb56","name":"Daft Punk"},{"id":"149f91ef-1287-46da-9a8e-87fee02f1471","name":"Pharrell Williams"}]}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 200, recorder.Code)
}

func TestRecordingHandler_WithNonExistingArtistIdAndExistingReleaseIdAndRecordingId(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/00000000-0000-0000-0000-000000000000/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings/833f00e1-781f-4edd-90e4-e52712618862.json")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"error":"artist not found"}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 404, recorder.Code)
}

func TestRecordingHandler_WithWrongArtistIdAndExistingReleaseIdAndRecordingId(t *testing.T) {
	// Artist is Guns'n'Roses but release is "Random Access Memories" by Daft Punk
	recorder := newTestRequest("GET", "/artists/eeb1195b-f213-4ce1-b28c-8565211f8e43/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings/833f00e1-781f-4edd-90e4-e52712618862.json")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"error":"release not found"}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 404, recorder.Code)
}

func TestRecordingHandler_WithExistingArtistIdAndReleaseIdAndRecordingId(t *testing.T) {
	recorder := newTestRequest("GET", "/artists/056e4f3e-d505-4dad-8ec1-d04f521cbb56/releases/79215cdf-4764-4dee-b0b9-fec1643df7c5/recordings/833f00e1-781f-4edd-90e4-e52712618862.json")

	body := string(recorder.Body.Bytes())
	expectedBody := `{"id":"833f00e1-781f-4edd-90e4-e52712618862","name":"Get Lucky","comment":"","length":367000,"artists":[{"id":"056e4f3e-d505-4dad-8ec1-d04f521cbb56","name":"Daft Punk"},{"id":"149f91ef-1287-46da-9a8e-87fee02f1471","name":"Pharrell Williams"}]}` + "\n"

	assert.Equal(t, expectedBody, body)
	assert.Equal(t, 200, recorder.Code)
}
