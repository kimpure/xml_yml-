import xml.etree.ElementTree as ET

tree = ET.parse('comp.xml');
root = tree.getroot();

def change_yml(element, indent=0):
    return_var = '' # 리턴할값 저장
    _ = '    ' * indent # yml 칸 나누는 크기
    if return_var == '':
        return_var = f'{element.tag}:\n    '
    else:
        return_var = f'{return_var}\n{_}{element.tag}:'
        
    for key in element:
        if len(key):
            return_var += change_yml(key, indent + 1)
        else:
            return_var = f'{return_var}\n{_}    {key.tag}: {key.text}'

    return return_var

yml_fp = 'comp.yml'
with open(yml_fp, 'w', encoding='utf-8') as f:
    f.write(change_yml(root))