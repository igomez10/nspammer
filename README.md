# nspammer

A Naive Bayes spam classifier implementation in Go, enabling text classification system using the Naive Bayes algorithm with Laplace smoothing to classify messages as spam or not spam.

## Features

- **Naive Bayes Classification**: Uses probabilistic classification based on Bayes' theorem with naive independence assumptions
- **Laplace Smoothing**: Implements additive smoothing to handle zero probabilities for unseen words
- **Training & Classification**: Simple API for training on labeled datasets and classifying new messages
- **Real Dataset Testing**: Includes tests with actual spam/ham email datasets

## Installation

```bash
go get github.com/igomez10/nspammer
```

## Usage

### Basic Example

```go
package main

import (
    "fmt"
    "github.com/igomez10/nspammer"
)

func main() {
    // Create training dataset (map[string]bool where true = spam, false = not spam)
    trainingData := map[string]bool{
        "buy viagra now":           true,
        "get rich quick":           true,
        "meeting at 3pm":           false,
        "project update report":    false,
    }

    // Create and train classifier
    classifier := nspammer.NewSpamClassifier(trainingData)

    // Classify new messages
    isSpam := classifier.Classify("buy now")
    fmt.Printf("Is spam: %v\n", isSpam)
}
```

### API

#### `NewSpamClassifier(dataset map[string]bool) *SpamClassifier`

Creates a new spam classifier and trains it on the provided dataset. The dataset is a map where keys are text messages and values indicate whether the message is spam (`true`) or not spam (`false`).

#### `(*SpamClassifier).Classify(input string) bool`

Classifies the input text as spam (`true`) or not spam (`false`) based on the trained model.

## How It Works

The classifier uses the Naive Bayes algorithm:

1. **Training Phase**:
   - Calculates prior probabilities: P(spam) and P(not spam)
   - Builds a vocabulary from all training messages
   - Counts word occurrences in spam and non-spam messages
   - Stores word frequencies for likelihood calculations

2. **Classification Phase**:
   - Calculates log probabilities to avoid numerical underflow
   - Computes: log(P(spam)) + Σ log(P(word|spam))
   - Computes: log(P(not spam)) + Σ log(P(word|not spam))
   - Returns `true` (spam) if the spam score is higher

3. **Laplace Smoothing**:
   - Adds a smoothing constant to avoid zero probabilities for unseen words
   - Formula: P(word|class) = (count + α) / (total + α × vocabulary_size)
   - Default α = 1.0

## Dataset

The project includes support for the Kaggle Spam Mails Dataset. To download it:

```bash
./init.sh
```

This script requires the [Kaggle CLI](https://github.com/Kaggle/kaggle-api) to be installed and configured.

## Testing

Run the test suite:

```bash
go test -v
```

The tests include:

- Simple classification examples
- Real-world email dataset evaluation
- Accuracy measurements on train/test splits
