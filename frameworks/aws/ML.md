# AWS Amazon ML

**Comprehend** - NLP to extract insights about the content of documents.

**Rekognition** - Add image and video analysis to your applications. Provide image/vid to the Rekognition API, servce can identify objects, people, text, scenes, and activities.

**Polly** - cloud service that converts text into lifelike speech.

**SageMaker** - fully managed machine learning service. Data scientists

Machine Learning web service offers _2 types of predictions_: 
- **Batch Predictions**: asynchronously generate predictions for multiple input data observations
- **Real-time Predictions**: synchronously generate predictions for individual data observations

### Binary Classification Model

Ex: Will customer buy this product? -> Yes, no (hence binary classification)

##### Training
To train **binary classification** models, Amazon ML uses industry-standard learning algorithim known as **logistic regression**.

Amazon ML provides industry-standard accuracy metric for **binary classification** models called **Area Under the (Receiver Operating Characteristic) Curve (AUC)**.

### Multiclass Classification Model

ML models for multiclass classifcation problems allow you to generate predictions for multiple classes. In Amazon ML, the **macro-average F1 score** is used to evaluate the predictive accuracy of a multiclass metric.


### Regression Model

ML models for regression problems predict a numeric value. Amazon ML uses the industry standard **Root Mean Square Error (RMSE)** metric.