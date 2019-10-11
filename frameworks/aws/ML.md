# AWS Amazon ML

Different types of logistic regression is implemented as Binary classification, and multiclass classification respectively and linear regression is implemeneted as regression model.

**Comprehend** - NLP to extract insights about the content of documents.

**Rekognition** - Add image and video analysis to your applications. Provide image/vid to the Rekognition API, servce can identify objects, people, text, scenes, and activities.

**Lex** - building conversational interfaces for applications using voice and text. Same engine that powers Alexa is now available to any developer, enabling natural language chatbots.

**Polly** - cloud service that converts text into lifelike speech. Use policy to develop apps that increase engagement and accessibility.

**SageMaker** - fully managed machine learning service. Data scientists can quickly and easily build and train machine learning models.

Machine Learning web service offers _2 types of predictions_: 
- **Batch Predictions**: asynchronously generate predictions for multiple input data observations
- **Real-time Predictions**: synchronously generate predictions for individual data observations

### Binary Classification Model

`Ex: Will customer buy this product? -> Yes / No` \
_(hence binary classification)_

To train **binary classification** models, Amazon ML uses industry-standard learning algorithim known as **logistic regression**.

`Test`: Amazon ML provides industry-standard accuracy metric for **binary classification** models called **Area Under the (Receiver Operating Characteristic) Curve (AUC)**.

### Multiclass Classification Model

(predict one of more than two outcomes)

ML models for multiclass classification problems allow you to generate predictions for **multiple classes**.

`Test`: Amazon ML uses the **macro-average F1 score** to evaluate the predictive accuracy of a multiclass metric.

### Regression Model

ML models for regression problems predict a **numeric value**. 

`Test`: Amazon ML uses the industry standard **Root Mean Square Error (RMSE)** metric.

### Hyperparameters / Overfitting
In Amazon Machine Learning, there are four hyperparameters that you can set: 
- number of passes
- regularization
- model size
- shuffle type

Prevent overfitting the models by using cross validation techniques.

**Overfitting** occurs when a model has memorized patterns that occur in the training and evaluation datasources, but has failed to generalize the patterns in the data.