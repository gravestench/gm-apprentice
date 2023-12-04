import pytesseract
import os
import json
import yaml

from PIL import Image

# Install the pytesseract library using pip if not already installed:
# pip install pytesseract

# Install Tesseract OCR and set the Tesseract executable path
# You can download Tesseract from: https://github.com/tesseract-ocr/tesseract
# Then, set the path to the Tesseract executable below:
pytesseract.pytesseract.tesseract_cmd = 'C:\\Program Files\\Tesseract-OCR\\tesseract.exe'

origins = [
    (57, 33),
    (412, 33),
    (767, 33),
    (57, 441),
    (412, 441),
    (767, 441),
]

rect_dict = {
    'Event1': (30, 109, 66, 13),
    'Event2': (105, 109, 66, 13),
    'Event3': (178, 109, 66, 13),
    'SensoryHear': (42, 128, 202, 17),
    'SensorySee': (42, 145, 202, 17),
    'SensoryTouch': (42, 162, 202, 17),
    'SensorySmell': (42, 178, 202, 17),
    'Belongings': (31, 230, 137, 45),
    'Name1': (173, 230, 140, 17),
    'Name2': (173, 246, 140, 17),
    'Name3': (173, 261, 140, 17),
    'Catalyst': (28, 292, 140, 18),
    'Virtue': (173, 292, 140, 18),
    'Location': (28, 327, 140, 16),
    'Vice': (173, 327, 74, 16),
}

def extract_text_from_image(image_path, rect_dict):
    """
    Extracts text from specified regions in an image and stores them in a result dictionary.

    Args:
        image_path (str): The path to the image file.
        rect_dict (dict): A dictionary with string IDs as keys and corresponding sub-rectangle coordinates (x, y, w, h) as values.

    Returns:
        dict: A dictionary where keys are string IDs and values are arrays of extracted text.
    """
    result = {}
    try:
        # Open the image using PIL (Python Imaging Library)
        image = Image.open(image_path)

        for origin in origins:
            for key, coordinates in rect_dict.items():
                x1, y1, w, h = coordinates
                x1 += origin[0]
                y1 += origin[1]
                x2 = x1 + w
                y2 = y1 + h

                # Crop the image to the specified coordinates
                cropped_image = image.crop((x1, y1, x2, y2))

                # Use Tesseract OCR to extract text from the cropped region
                extracted_text = pytesseract.image_to_string(cropped_image)
                extracted_text = extracted_text.replace("\n", " ")
                extracted_text = extracted_text.strip("|\'\" \u201C  -—:;‘.~_)")
                extracted_text = extracted_text.replace("..", "...")
                if extracted_text == "":
                    continue

                # Store the extracted text in the result dictionary
                if key in result:
                    result[key].append(extracted_text)
                else:
                    result[key] = [extracted_text]

    except Exception as e:
        print(f"Error processing image: {str(e)}")

    return result

def unique_strings(input_array):
    """
    Returns an array of unique strings from the input array.

    Args:
        input_array (list): The input array of strings.

    Returns:
        list: An array containing unique strings from the input array.
    """
    unique_set = set(input_array)
    unique_list = list(unique_set)
    return unique_list

# Specify the directory containing PNG files
input_directory = 'C:\\Users\\dknut\\src\\dm-apprentice\\cmd\\scrape\\data'

# List all PNG files in the directory
png_files = [f for f in os.listdir(input_directory) if f.endswith('.png')]

result = {}

# Iterate through the PNG files and process each one
for png_file in png_files:
    image_path = os.path.join(input_directory, png_file)

    extracted = extract_text_from_image(image_path, rect_dict)

    for key, texts in extracted.items():
        if key in result:
            result[key].append(texts)
        else:
            result[key] = texts
        try:
            result[key] = unique_strings(result[key])
        except Exception as e:
            continue

def encode_dict_to_json(input_dict):
    """
    Encodes a dictionary into JSON format.

    Args:
        input_dict (dict): The dictionary to be encoded.

    Returns:
        str: A JSON-encoded string representing the input dictionary.
    """
    encoded_json = json.dumps(input_dict, indent=4)  # Use indent for pretty formatting
    return encoded_json

def encode_dict_to_yaml(input_dict):
    """
    Encodes a dictionary into YAML format.

    Args:
        input_dict (dict): The dictionary to be encoded.

    Returns:
        str: A YAML-encoded string representing the input dictionary.
    """
    encoded_yaml = yaml.dump(input_dict, default_flow_style=False)
    return encoded_yaml

def write_string_to_file(filename, content):
    """
    Writes a string to a file.

    Args:
        filename (str): The name of the file to write to.
        content (str): The string content to write to the file.
    """
    with open(filename, 'w') as file:
        file.write(content)

write_string_to_file("scraped.yaml", encode_dict_to_yaml(result))