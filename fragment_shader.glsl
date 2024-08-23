#version 410 core
precision highp float;  // Ensure high precision

out vec4 FragColor;

in vec3 FragPos;
in vec3 Normal;

struct Light {
    vec3 position;
    vec3 color;
    float intensity;
};

uniform int numLights;
uniform Light lights[10];
uniform vec3 viewPos;
uniform vec3 objectColor;

void main()
{
    vec3 result = vec3(0.0);
    vec3 norm = normalize(Normal);
    vec3 viewDir = normalize(viewPos - FragPos);

    for (int i = 0; i < numLights; i++) {
        // Ambient
        float ambientStrength = 0.1;
        vec3 ambient = ambientStrength * lights[i].color * lights[i].intensity;

        // Diffuse
        vec3 lightDir = normalize(lights[i].position - FragPos);
        float diff = max(dot(norm, lightDir), 0.0);
        vec3 diffuse = diff * lights[i].color * lights[i].intensity;

        // Specular
        float specularStrength = 0.5;
        vec3 reflectDir = reflect(-lightDir, norm);
        float spec = pow(max(dot(viewDir, reflectDir), 0.0), 32);  // The '32' controls the shininess
        vec3 specular = specularStrength * spec * lights[i].color * lights[i].intensity;

        result += (ambient + diffuse + specular);
    }

    FragColor = vec4(result * objectColor, 1.0);
}
